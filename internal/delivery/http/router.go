// Package http nối route Gin cho AMF Event Exposure Service.
package http

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/KhanhCt-study/amf-event-exposure.git/internal/delivery/http/api"
	"github.com/KhanhCt-study/amf-event-exposure.git/internal/delivery/http/event"
	subhandler "github.com/KhanhCt-study/amf-event-exposure.git/internal/delivery/http/subscription"
	"github.com/KhanhCt-study/amf-event-exposure.git/pkg/http/middleware"
)

// BasePath là base path của service (TS 29.518 §6.2.1).
const BasePath = "/namf-evts/v1"

// NewRouter dựng gin.Engine với recovery, logger và các route subscription.
func NewRouter(handler *subhandler.AmfSubHandler, trigger *event.TriggerHandler, logger *slog.Logger) *gin.Engine {
	if logger == nil {
		logger = slog.Default()
	}

	engine := gin.New()
	engine.RedirectTrailingSlash = false
	engine.HandleMethodNotAllowed = true
	engine.Use(gin.Logger(), middleware.Recovery(logger))

	engine.NoRoute(func(c *gin.Context) {
		api.WriteError(c, http.StatusNotFound, "resource not found")
	})
	engine.NoMethod(func(c *gin.Context) {
		api.WriteError(c, http.StatusMethodNotAllowed, "method not allowed for this resource")
	})

	engine.GET("/healthz", func(c *gin.Context) {
		api.WriteSuccess(c, gin.H{"status": "ok"})
	})

	group := engine.Group(BasePath)
	{
		group.POST("/subscriptions", handler.CreateSubscription)
		group.PATCH("/subscriptions/:"+subhandler.ParamSubscriptionID, handler.ModifySubscription)
		group.DELETE("/subscriptions/:"+subhandler.ParamSubscriptionID, handler.DeleteSubscription)
	}

	// Route ngoài 3GPP: kích hoạt thủ công một event report để test luồng
	// notification khi chưa có UE/NGAP thật.
	if trigger != nil {
		engine.POST("/internal/v1/event-reports", trigger.TriggerReport)
	}
	return engine
}
