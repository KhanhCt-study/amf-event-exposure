package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KhanhCt-study/amf-event-exposure.git/internal/domain"
	model "github.com/KhanhCt-study/amf-event-exposure.git/pkg/3gppmodel"
)

// StatusFromDomainError: ErrInvalid→400, ErrUnsupported→403, ErrNotFound→404,
// ErrConflict→409, còn lại→500.
func StatusFromDomainError(err error) int {
	switch {
	case errors.Is(err, domain.ErrInvalid):
		return http.StatusBadRequest
	case errors.Is(err, domain.ErrUnsupported):
		return http.StatusForbidden
	case errors.Is(err, domain.ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, domain.ErrConflict):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func stringPtr(s string) *string {
	return &s
}

// ProblemFromDomainError dựng ProblemDetails, kèm invalidParams khi lỗi mang
// thông tin field (FieldError / FieldErrors).
func ProblemFromDomainError(err error) *model.ProblemDetails {
	status := StatusFromDomainError(err)
	problem := &model.ProblemDetails{
		Status: &status,
		Title:  stringPtr(http.StatusText(status)),
		Detail: stringPtr(err.Error()),
	}

	if status == http.StatusInternalServerError {
		problem.Detail = stringPtr("internal server error")
		return problem
	}

	var opErr *domain.OperationError
	if errors.As(err, &opErr) {
		problem.Detail = stringPtr(opErr.Err.Error())
	}

	var fieldErrs domain.FieldErrors
	if errors.As(err, &fieldErrs) {
		var params []model.InvalidParam
		if problem.InvalidParams != nil {
			params = *problem.InvalidParams
		}
		for _, fe := range fieldErrs {
			params = append(params, model.InvalidParam{
				Param:  fe.Field,
				Reason: &fe.Reason,
			})
		}
		if len(params) > 0 {
			problem.InvalidParams = &params
		}
		return problem
	}

	var fieldErr *domain.FieldError
	if errors.As(err, &fieldErr) && fieldErr.Field != "" {
		problem.InvalidParams = &[]model.InvalidParam{{Param: fieldErr.Field, Reason: &fieldErr.Reason}}
	}
	return problem
}

// WriteDomainError map error nghiệp vụ sang response ProblemDetails tương ứng.
func WriteDomainError(c *gin.Context, err error) {
	WriteProblem(c, ProblemFromDomainError(err))
}
