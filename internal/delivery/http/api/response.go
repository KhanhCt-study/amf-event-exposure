package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	model "github.com/KhanhCt-study/amf-event-exposure.git/pkg/3gppmodel"
)

// writeJSON ghi resource dưới dạng JSON thô (KHÔNG bọc envelope) cho các route
// 3GPP. Với HEAD, chỉ ghi header và bỏ qua body.
func writeJSON(c *gin.Context, status int, contentType string, resource interface{}) {
	if resource == nil {
		c.Status(status)
		c.Writer.WriteHeaderNow()
		return
	}

	body, err := json.Marshal(resource)
	if err != nil {
		c.Header("Content-Type", ContentTypeProblem)
		c.Status(http.StatusInternalServerError)
		_, _ = c.Writer.Write(mustMarshalProblem(http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError), "cannot encode response body"))
		return
	}

	c.Header("Content-Type", contentType)
	if c.Request != nil && c.Request.Method == http.MethodHead {
		c.Status(status)
		c.Writer.WriteHeaderNow()
		return
	}
	c.Status(status)
	_, _ = c.Writer.Write(body)
}

// WriteCreated ghi 201 kèm Location header (URI đầy đủ của resource).
func WriteCreated(c *gin.Context, location string, resource interface{}) {
	if location != "" {
		c.Header("Location", location)
	}
	writeJSON(c, http.StatusCreated, ContentTypeJSON, resource)
}

// WriteSuccess ghi 200 với resource.
func WriteSuccess(c *gin.Context, resource interface{}) {
	writeJSON(c, http.StatusOK, ContentTypeJSON, resource)
}

// WriteNoContent ghi 204, không body.
func WriteNoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
	c.Writer.WriteHeaderNow()
}

// WriteError ghi ProblemDetails với status và detail chỉ định.
func WriteError(c *gin.Context, status int, detail string) {
	WriteProblem(c, &model.ProblemDetails{
		Status: &status,
		Title:  stringPtr(http.StatusText(status)),
		Detail: stringPtr(detail),
	})
}

// WriteProblem ghi một ProblemDetails đã dựng sẵn (application/problem+json).
func WriteProblem(c *gin.Context, problem *model.ProblemDetails) {
	status := int(*problem.Status)
	if status == 0 {
		status = http.StatusInternalServerError
		problem.Status = &status
	}
	if problem.Title == stringPtr("") {
		problem.Title = stringPtr(http.StatusText(status))
	}
	if problem.Instance == stringPtr("") && c.Request != nil {
		problem.Instance = &c.Request.URL.Path
	}
	writeJSON(c, status, ContentTypeProblem, problem)
}

func mustMarshalProblem(status int, title, detail string) []byte {
	body, _ := json.Marshal(&model.ProblemDetails{
		Status: &status, Title: &title, Detail: &detail,
	})
	return body
}
