package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/KhanhCt-study/amf-event-exposure.git/internal/domain"
	model "github.com/KhanhCt-study/amf-event-exposure.git/pkg/3gppmodel"
)

func newTestContext(method, target string) (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = httptest.NewRequest(method, target, nil)
	return c, rec
}

func TestWriteCreated_SetsLocationAndBody(t *testing.T) {
	c, rec := newTestContext(http.MethodPost, "/namf-evts/v1/subscriptions")
	location := "http://amf:8081/namf-evts/v1/subscriptions/abc"

	WriteCreated(c, location, &model.AmfCreatedEventSubscription{SubscriptionId: location})

	if rec.Code != http.StatusCreated {
		t.Errorf("status = %d, want 201", rec.Code)
	}
	if got := rec.Header().Get("Location"); got != location {
		t.Errorf("Location = %q, want %q", got, location)
	}
	if got := rec.Header().Get("Content-Type"); got != ContentTypeJSON {
		t.Errorf("Content-Type = %q, want %q", got, ContentTypeJSON)
	}

	var body model.AmfCreatedEventSubscription
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("response không phải JSON hợp lệ: %v", err)
	}
	if body.SubscriptionId != location {
		t.Errorf("subscriptionId = %q, want %q", body.SubscriptionId, location)
	}
}

func TestWriteSuccess_HeadSkipsBody(t *testing.T) {
	c, rec := newTestContext(http.MethodHead, "/namf-evts/v1/subscriptions/abc")

	WriteSuccess(c, map[string]string{"hello": "world"})

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", rec.Code)
	}
	if rec.Body.Len() != 0 {
		t.Errorf("HEAD không được ghi body, got %q", rec.Body.String())
	}
}

func TestWriteNoContent(t *testing.T) {
	c, rec := newTestContext(http.MethodDelete, "/namf-evts/v1/subscriptions/abc")

	WriteNoContent(c)

	if rec.Code != http.StatusNoContent {
		t.Errorf("status = %d, want 204", rec.Code)
	}
	if rec.Body.Len() != 0 {
		t.Errorf("204 không được có body, got %q", rec.Body.String())
	}
}

func TestWriteDomainError_Mapping(t *testing.T) {
	tests := []struct {
		name       string
		err        error
		wantStatus int
	}{
		{"invalid", fmt.Errorf("%w: bad", domain.ErrInvalid), http.StatusBadRequest},
		{"unsupported", fmt.Errorf("%w: nope", domain.ErrUnsupported), http.StatusForbidden},
		{"not found", fmt.Errorf("%w: gone", domain.ErrNotFound), http.StatusNotFound},
		{"conflict", fmt.Errorf("%w: race", domain.ErrConflict), http.StatusConflict},
		{"unknown", fmt.Errorf("boom"), http.StatusInternalServerError},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c, rec := newTestContext(http.MethodPost, "/namf-evts/v1/subscriptions")

			WriteDomainError(c, tc.err)

			if rec.Code != tc.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tc.wantStatus)
			}
			if got := rec.Header().Get("Content-Type"); got != ContentTypeProblem {
				t.Errorf("Content-Type = %q, want %q", got, ContentTypeProblem)
			}
		})
	}
}

func TestWriteDomainError_InvalidParams(t *testing.T) {
	c, rec := newTestContext(http.MethodPost, "/namf-evts/v1/subscriptions")
	errs := domain.FieldErrors{
		domain.NewFieldError(domain.ErrInvalid, "/subscription/eventList", "at least one event is required"),
		domain.NewFieldError(domain.ErrInvalid, "/subscription/nfId", "must be a valid UUID"),
	}

	WriteDomainError(c, errs)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
	var problem model.ProblemDetails
	if err := json.Unmarshal(rec.Body.Bytes(), &problem); err != nil {
		t.Fatalf("problem không phải JSON hợp lệ: %v", err)
	}
	if len(*problem.InvalidParams) != 2 {
		t.Fatalf("invalidParams = %d, want 2", len(*problem.InvalidParams))
	}
	if (*problem.InvalidParams)[0].Param != "/subscription/eventList" {
		t.Errorf("param đầu tiên = %q", (*problem.InvalidParams)[0].Param)
	}
}
