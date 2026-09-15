// Package event cung cấp hook nội bộ (KHÔNG thuộc TS 29.518) để kích hoạt
// event report thủ công trong môi trường prototype/test.
package event

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/KhanhCt-study/amf-event-exposure.git/internal/delivery/http/api"
	"github.com/KhanhCt-study/amf-event-exposure.git/internal/domain"
	"github.com/KhanhCt-study/amf-event-exposure.git/internal/domain/interfaces"
	model "github.com/KhanhCt-study/amf-event-exposure.git/pkg/3gppmodel"
)

// TriggerRequest là payload của hook nội bộ.
type TriggerRequest struct {
	SubscriptionID string              `json:"subscriptionId"`
	Type           model.AmfEventType  `json:"type"`
	Supi           string              `json:"supi,omitempty"`
	Location       *model.UserLocation `json:"location,omitempty"`
}

// TriggerHandler đẩy report vào notification usecase.
type TriggerHandler struct {
	notifier interfaces.INotificationUseCase
}

// NewTriggerHandler dựng handler với constructor injection.
func NewTriggerHandler(notifier interfaces.INotificationUseCase) (*TriggerHandler, error) {
	if notifier == nil {
		return nil, errors.New("event: notifier must not be nil")
	}
	return &TriggerHandler{notifier: notifier}, nil
}

// TriggerReport — POST /internal/v1/event-reports → 204.
func (h *TriggerHandler) TriggerReport(c *gin.Context) {
	var req TriggerRequest
	if err := api.Decode(c, &req); err != nil {
		api.WriteDomainError(c, err)
		return
	}
	if req.SubscriptionID == "" {
		api.WriteDomainError(c, domain.NewFieldError(domain.ErrInvalid, "/subscriptionId", "subscriptionId is required"))
		return
	}
	if req.Type == "" {
		req.Type = model.AmfEventTypeLocationReport
	}

	report := model.NamfEventReport{
		Type:      req.Type,
		TimeStamp: time.Now().UTC().Format(time.RFC3339),
		Supi:      req.Supi,
		Location:  req.Location,
	}
	if err := h.notifier.DispatchReport(c.Request.Context(), req.SubscriptionID, report); err != nil {
		api.WriteDomainError(c, err)
		return
	}
	api.WriteNoContent(c)
}
