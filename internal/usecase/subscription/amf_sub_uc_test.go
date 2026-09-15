package subscription

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/KhanhCt-study/amf-event-exposure.git/internal/domain"
	gppmodel "github.com/KhanhCt-study/amf-event-exposure.git/pkg/3gppmodel"
	model "github.com/KhanhCt-study/amf-event-exposure.git/pkg/3gppmodel"
	"github.com/google/uuid"
)

// --- manual mocks (không dùng mockgen) ---

type mockRepo struct {
	subs        map[string]*domain.AmfSubscription
	createErr   error
	updateErr   error
	deleteErr   error
	deleteCalls []string
}

func newMockRepo() *mockRepo {
	return &mockRepo{subs: make(map[string]*domain.AmfSubscription)}
}

func (m *mockRepo) Create(_ context.Context, sub *domain.AmfSubscription) error {
	if m.createErr != nil {
		return m.createErr
	}
	copied := *sub
	m.subs[sub.ID] = &copied
	return nil
}

func (m *mockRepo) GetByID(_ context.Context, id string) (*domain.AmfSubscription, error) {
	sub, ok := m.subs[id]
	if !ok {
		return nil, fmt.Errorf("%w: subscription %s", domain.ErrNotFound, id)
	}
	copied := *sub
	return &copied, nil
}

func (m *mockRepo) ListActiveBySupi(_ context.Context, _ string) ([]*domain.AmfSubscription, error) {
	return nil, nil
}

func (m *mockRepo) Update(_ context.Context, sub *domain.AmfSubscription) error {
	if m.updateErr != nil {
		return m.updateErr
	}
	current, ok := m.subs[sub.ID]
	if !ok {
		return fmt.Errorf("%w: subscription %s", domain.ErrNotFound, sub.ID)
	}
	if current.Version != sub.Version {
		return fmt.Errorf("%w: subscription %s", domain.ErrConflict, sub.ID)
	}
	sub.Version++
	copied := *sub
	m.subs[sub.ID] = &copied
	return nil
}

func (m *mockRepo) Delete(_ context.Context, id string, expectedVersion uint64) error {
	m.deleteCalls = append(m.deleteCalls, id)
	if m.deleteErr != nil {
		return m.deleteErr
	}
	current, ok := m.subs[id]
	if !ok {
		return fmt.Errorf("%w: subscription %s", domain.ErrNotFound, id)
	}
	if current.Version != expectedVersion {
		return fmt.Errorf("%w: subscription %s", domain.ErrConflict, id)
	}
	delete(m.subs, id)
	return nil
}

type mockLocations struct{ err error }

func (m *mockLocations) Location(_ context.Context, _ string) (*model.UserLocation, error) {
	if m.err != nil {
		return nil, m.err
	}
	return &model.UserLocation{NrLocation: &model.NrLocation{
		Tai:  model.Tai{PlmnId: model.PlmnId{Mcc: "208", Mnc: "95"}, Tac: "000001"},
		Ncgi: model.Ncgi{PlmnId: model.PlmnId{Mcc: "208", Mnc: "95"}, NrCellId: "000000001"},
	}}, nil
}

// --- helpers ---

const fixedID = "a1b2c3d4-0000-4000-8000-000000000001"

func newTestUseCase(t *testing.T, repo *mockRepo) *AmfSubUseCaseImpl {
	t.Helper()
	uc, err := NewAmfSubUseCase(repo, &mockLocations{}, "http://amf:8081")
	if err != nil {
		t.Fatalf("NewAmfSubUseCase: %v", err)
	}
	uc.newID = func() string { return fixedID }
	uc.now = func() time.Time { return time.Date(2026, 9, 15, 10, 0, 0, 0, time.UTC) }
	return uc
}

func validRequest() *model.AmfCreateEventSubscription {
	return &model.AmfCreateEventSubscription{
		Subscription: model.AmfEventSubscription{
			EventList: []model.AmfEvent{
				{Type: model.AmfEventTypeLocationReport, ImmediateFlag: true},
			},
			EventNotifyUri:      "http://nwdaf:8080/callbacks/amf-notify",
			NotifyCorrelationId: "corr-abc-123",
			NfId:                gppmodel.NfInstanceId(uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")),
			Supi:                "imsi-208950000000010",
			Options:             &model.AmfEventMode{Trigger: model.AmfEventTriggerContinuous},
		},
	}
}

// --- tests ---

func TestCreateSubscription_Success(t *testing.T) {
	repo := newMockRepo()
	uc := newTestUseCase(t, repo)

	created, err := uc.CreateSubscription(context.Background(), validRequest())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	wantURI := "http://amf:8081/namf-evts/v1/subscriptions/" + fixedID
	if created.SubscriptionId != wantURI {
		t.Errorf("subscriptionId = %q, want %q", created.SubscriptionId, wantURI)
	}
	if len(created.ReportList) != 1 {
		t.Fatalf("immediateFlag=true phải trả 1 report, got %d", len(created.ReportList))
	}
	if created.ReportList[0].Location == nil {
		t.Error("report thiếu location")
	}
	if _, ok := repo.subs[fixedID]; !ok {
		t.Error("subscription chưa được lưu vào repository")
	}
}

func TestCreateSubscription_ValidationErrors(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(*model.AmfCreateEventSubscription)
		wantErr error
	}{
		{"eventList rỗng", func(r *model.AmfCreateEventSubscription) {
			r.Subscription.EventList = nil
		}, domain.ErrInvalid},
		{"eventNotifyUri không phải URL tuyệt đối", func(r *model.AmfCreateEventSubscription) {
			r.Subscription.EventNotifyUri = "/callbacks/amf-notify"
		}, domain.ErrInvalid},
		{"nfId không phải UUID", func(r *model.AmfCreateEventSubscription) {
			r.Subscription.NfId = gppmodel.NfInstanceId(uuid.Nil)
		}, domain.ErrInvalid},
		{"supi sai pattern", func(r *model.AmfCreateEventSubscription) {
			r.Subscription.Supi = "208950000000010"
		}, domain.ErrInvalid},
		{"event type ngoài prototype", func(r *model.AmfCreateEventSubscription) {
			r.Subscription.EventList[0].Type = model.AmfEventTypeReachabilityReport
		}, domain.ErrUnsupported},
		{"trigger không hợp lệ", func(r *model.AmfCreateEventSubscription) {
			r.Subscription.Options.Trigger = "HOURLY"
		}, domain.ErrInvalid},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			repo := newMockRepo()
			uc := newTestUseCase(t, repo)
			req := validRequest()
			tc.mutate(req)

			_, err := uc.CreateSubscription(context.Background(), req)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("err = %v, want errors.Is(..., %v)", err, tc.wantErr)
			}
			if len(repo.subs) != 0 {
				t.Error("không được lưu subscription khi validate thất bại")
			}
		})
	}
}

func TestModifySubscription_ReplaceEvent(t *testing.T) {
	repo := newMockRepo()
	uc := newTestUseCase(t, repo)
	if _, err := uc.CreateSubscription(context.Background(), validRequest()); err != nil {
		t.Fatalf("setup: %v", err)
	}

	maxReports := int32(50)
	updated, err := uc.ModifySubscription(context.Background(), fixedID,
		[]model.AmfUpdateEventSubscriptionItem{{
			Op:   "replace",
			Path: "/subscription/eventList/0",
			Value: &model.AmfEvent{
				Type:          model.AmfEventTypeLocationReport,
				ImmediateFlag: false,
				MaxReports:    int(maxReports),
			},
		}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	ev := updated.Subscription.EventList[0]
	if ev.ImmediateFlag {
		t.Error("immediateFlag phải là false sau khi patch")
	}
	if ev.MaxReports == 0 || ev.MaxReports != int(maxReports) {
		t.Errorf("maxReports = %v, want %d", ev.MaxReports, maxReports)
	}
	if got := repo.subs[fixedID].Version; got != 2 {
		t.Errorf("version = %d, want 2 (CAS tăng version)", got)
	}
}

func TestModifySubscription_UnsupportedPath(t *testing.T) {
	repo := newMockRepo()
	uc := newTestUseCase(t, repo)
	if _, err := uc.CreateSubscription(context.Background(), validRequest()); err != nil {
		t.Fatalf("setup: %v", err)
	}

	_, err := uc.ModifySubscription(context.Background(), fixedID,
		[]model.AmfUpdateEventSubscriptionItem{{
			Op:    "replace",
			Path:  "/subscription/nfId",
			Value: &model.AmfEvent{},
		}})
	if !errors.Is(err, domain.ErrUnsupported) {
		t.Fatalf("err = %v, want ErrUnsupported", err)
	}
}

func TestModifySubscription_NotFound(t *testing.T) {
	uc := newTestUseCase(t, newMockRepo())

	_, err := uc.ModifySubscription(context.Background(), fixedID,
		[]model.AmfUpdateEventSubscriptionItem{{
			Op:    "replace",
			Path:  "/subscription/options",
			Value: &model.AmfEvent{},
		}})
	if !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("err = %v, want ErrNotFound", err)
	}
}

func TestDeleteSubscription(t *testing.T) {
	repo := newMockRepo()
	uc := newTestUseCase(t, repo)
	if _, err := uc.CreateSubscription(context.Background(), validRequest()); err != nil {
		t.Fatalf("setup: %v", err)
	}

	if err := uc.DeleteSubscription(context.Background(), fixedID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(repo.subs) != 0 {
		t.Error("subscription vẫn còn sau khi xóa")
	}

	err := uc.DeleteSubscription(context.Background(), fixedID)
	if !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("xóa lần hai: err = %v, want ErrNotFound", err)
	}
}

func TestDeleteSubscription_InvalidID(t *testing.T) {
	uc := newTestUseCase(t, newMockRepo())
	if err := uc.DeleteSubscription(context.Background(), "abc"); !errors.Is(err, domain.ErrInvalid) {
		t.Fatalf("err = %v, want ErrInvalid", err)
	}
}
