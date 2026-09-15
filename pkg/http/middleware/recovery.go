// Package middleware chứa các Gin middleware dùng chung.
package middleware

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Recovery bắt panic, log lại và trả 500 dạng application/problem+json
// thay vì HTML mặc định của Gin.
func Recovery(logger *slog.Logger) gin.HandlerFunc {
	if logger == nil {
		logger = slog.Default()
	}
	return func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				logger.Error("panic recovered",
					"method", c.Request.Method, "path", c.Request.URL.Path, "panic", rec)
				c.Header("Content-Type", "application/problem+json")
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"status": http.StatusInternalServerError,
					"title":  http.StatusText(http.StatusInternalServerError),
					"detail": "internal server error",
				})
			}
		}()
		c.Next()
	}
}
