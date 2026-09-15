// Package api gom các helper dùng chung cho mọi HTTP handler:
// decode request, ghi response và map domain error sang HTTP status.
package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/KhanhCt-study/amf-event-exposure.git/internal/domain"
)

// MaxBodyBytes giới hạn kích thước request body (1 MiB).
const MaxBodyBytes int64 = 1 << 20

const (
	ContentTypeJSON      = "application/json"
	ContentTypeJSONPatch = "application/json-patch+json"
	ContentTypeProblem   = "application/problem+json"
)

// Decode đọc body JSON vào dst với DisallowUnknownFields và giới hạn 1 MiB.
// Mọi lỗi trả về đều wrap domain.ErrInvalid nên handler chỉ cần
// WriteDomainError là ra đúng 400 + ProblemDetails.
func Decode(c *gin.Context, dst interface{}, allowedContentTypes ...string) error {
	if len(allowedContentTypes) == 0 {
		allowedContentTypes = []string{ContentTypeJSON}
	}
	if err := checkContentType(c.GetHeader("Content-Type"), allowedContentTypes); err != nil {
		return err
	}

	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)
	dec := json.NewDecoder(c.Request.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		return decodeError(err)
	}
	// Không cho phép dữ liệu thừa sau JSON value đầu tiên.
	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return domain.NewFieldError(domain.ErrInvalid, "", "body must contain a single JSON value")
	}
	return nil
}

func checkContentType(header string, allowed []string) error {
	if header == "" {
		return domain.NewFieldError(domain.ErrInvalid, "", "Content-Type header is required")
	}
	mediaType, _, err := mime.ParseMediaType(header)
	if err != nil {
		return domain.NewFieldError(domain.ErrInvalid, "", "malformed Content-Type header")
	}
	for _, a := range allowed {
		if strings.EqualFold(mediaType, a) {
			return nil
		}
	}
	return domain.NewFieldError(domain.ErrInvalid, "",
		fmt.Sprintf("unsupported Content-Type %q, expected %s", mediaType, strings.Join(allowed, " or ")))
}

func decodeError(err error) error {
	var maxErr *http.MaxBytesError
	if errors.As(err, &maxErr) {
		return domain.NewFieldError(domain.ErrInvalid, "", "request body exceeds 1 MiB limit")
	}
	if errors.Is(err, io.EOF) {
		return domain.NewFieldError(domain.ErrInvalid, "", "request body is empty")
	}

	var typeErr *json.UnmarshalTypeError
	if errors.As(err, &typeErr) {
		field := typeErr.Field
		if field == "" {
			field = "body"
		}
		return domain.NewFieldError(domain.ErrInvalid, "/"+strings.ReplaceAll(field, ".", "/"),
			fmt.Sprintf("expected type %s", typeErr.Type.String()))
	}

	var syntaxErr *json.SyntaxError
	if errors.As(err, &syntaxErr) {
		return domain.NewFieldError(domain.ErrInvalid, "",
			fmt.Sprintf("malformed JSON at offset %d", syntaxErr.Offset))
	}

	if msg := err.Error(); strings.Contains(msg, "unknown field") {
		return domain.NewFieldError(domain.ErrInvalid, "", msg)
	}
	return domain.NewFieldError(domain.ErrInvalid, "", "cannot decode request body")
}
