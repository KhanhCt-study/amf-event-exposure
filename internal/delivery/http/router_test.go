package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	subhandler "github.com/KhanhCt-study/amf-event-exposure.git/internal/delivery/http/subscription"
	model "github.com/KhanhCt-study/amf-event-exposure.git/pkg/3gppmodel"
)

// stubUseCase panic khi được gọi — dùng để kiểm tra recovery middleware.
type stubUseCase struct{ panicOnCreate bool }

func (s *stubUseCase) CreateSubscription(
	_ context.Context, _ *model.AmfCreateEventSubscription,
) (*model.AmfCreatedEventSubscription, error) {
	if s.panicOnCreate {
		panic("boom")
	}
	return &model.AmfCreatedEventSubscription{}, nil
}

func (s *stubUseCase) ModifySubscription(
	_ context.Context, _ string, _ []model.AmfUpdateEventSubscriptionItem,
) (*model.AmfUpdatedEventSubscription, error) {
	return &model.AmfUpdatedEventSubscription{}, nil
}

func (s *stubUseCase) DeleteSubscription(_ context.Context, _ string) error { return nil }

func newTestRouter(t *testing.T, uc *stubUseCase) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)
	handler, err := subhandler.NewAmfSubHandler(uc)
	if err != nil {
		t.Fatalf("NewAmfSubHandler: %v", err)
	}
	return NewRouter(handler, nil, nil)
}

func TestRouter_NotFound(t *testing.T) {
	router := newTestRouter(t, &stubUseCase{})
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/namf-evts/v1/unknown", nil))

	if rec.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404", rec.Code)
	}
	if got := rec.Header().Get("Content-Type"); got != "application/problem+json" {
		t.Errorf("Content-Type = %q", got)
	}
}

func TestRouter_MethodNotAllowed(t *testing.T) {
	router := newTestRouter(t, &stubUseCase{})
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/namf-evts/v1/subscriptions", nil))

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("status = %d, want 405", rec.Code)
	}
}

func TestRouter_PanicRecovered(t *testing.T) {
	router := newTestRouter(t, &stubUseCase{panicOnCreate: true})
	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodPost, "/namf-evts/v1/subscriptions",
		strings.NewReader(`{"subscription":{"eventList":[]}}`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("status = %d, want 500", rec.Code)
	}
	if got := rec.Header().Get("Content-Type"); !strings.HasPrefix(got, "application/problem+json") {
		t.Errorf("Content-Type = %q", got)
	}
}

func TestRouter_RejectsMissingContentType(t *testing.T) {
	router := newTestRouter(t, &stubUseCase{})
	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodPost, "/namf-evts/v1/subscriptions",
		strings.NewReader(`{"subscription":{}}`))
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", rec.Code)
	}
}

func TestRouter_RejectsUnknownFields(t *testing.T) {
	router := newTestRouter(t, &stubUseCase{})
	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodPost, "/namf-evts/v1/subscriptions",
		strings.NewReader(`{"subscription":{"eventList":[]},"unexpected":1}`))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400 (DisallowUnknownFields)", rec.Code)
	}
}
