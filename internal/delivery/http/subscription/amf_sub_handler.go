// Package subscription chứa HTTP handler cho resource
// /namf-evts/v1/subscriptions (TS 29.518 §6.2.3).
package subscription

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/KhanhCt-study/amf-event-exposure.git/internal/delivery/http/api"
	"github.com/KhanhCt-study/amf-event-exposure.git/internal/domain/interfaces"
	model "github.com/KhanhCt-study/amf-event-exposure.git/pkg/3gppmodel"
)

// ParamSubscriptionID là tên path param của individual subscription.
const ParamSubscriptionID = "subscriptionId"

// AmfSubHandler map HTTP request sang IAmfEventSubUseCase.
type AmfSubHandler struct {
	useCase interfaces.IAmfEventSubUseCase
}

// NewAmfSubHandler dựng handler với constructor injection.
func NewAmfSubHandler(useCase interfaces.IAmfEventSubUseCase) (*AmfSubHandler, error) {
	if useCase == nil {
		return nil, errors.New("subscription: use case must not be nil")
	}
	return &AmfSubHandler{useCase: useCase}, nil
}

// CreateSubscription — POST /namf-evts/v1/subscriptions → 201 + Location.
func (h *AmfSubHandler) CreateSubscription(c *gin.Context) {
	var req model.AmfCreateEventSubscription
	if err := api.Decode(c, &req); err != nil {
		api.WriteDomainError(c, err)
		return
	}

	created, err := h.useCase.CreateSubscription(c.Request.Context(), &req)
	if err != nil {
		api.WriteDomainError(c, err)
		return
	}
	api.WriteCreated(c, created.SubscriptionId, created)
}

// ModifySubscription — PATCH /namf-evts/v1/subscriptions/{id} → 200.
func (h *AmfSubHandler) ModifySubscription(c *gin.Context) {
	var patches []model.AmfUpdateEventSubscriptionItem
	if err := api.Decode(c, &patches, api.ContentTypeJSON, api.ContentTypeJSONPatch); err != nil {
		api.WriteDomainError(c, err)
		return
	}

	updated, err := h.useCase.ModifySubscription(c.Request.Context(), c.Param(ParamSubscriptionID), patches)
	if err != nil {
		api.WriteDomainError(c, err)
		return
	}
	api.WriteSuccess(c, updated)
}

// DeleteSubscription — DELETE /namf-evts/v1/subscriptions/{id} → 204.
func (h *AmfSubHandler) DeleteSubscription(c *gin.Context) {
	if err := h.useCase.DeleteSubscription(c.Request.Context(), c.Param(ParamSubscriptionID)); err != nil {
		api.WriteDomainError(c, err)
		return
	}
	api.WriteNoContent(c)
}
