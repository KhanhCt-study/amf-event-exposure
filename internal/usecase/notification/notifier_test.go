package notification

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

type mockRepo struct {
	sub       *domain.AmfSubscription
	updateErr error
	updates   int
}

func (m *mockRepo) Create(context.Context, *domain.AmfSubscription) error { return nil }

func (m *mockRepo) GetByID(_ context.Context, id string) (*domain.AmfSubscription, error) {
	if m.sub == nil || m.sub.ID != id {
		return nil, fmt.Errorf("%w: subscription %s", domain.ErrNotFound, id)
	}
	copied := *m.sub
	if m.sub.RemainReports != nil {
		remain := *m.sub.RemainReports
		copied.RemainReports = &remain
	}
	return &copied, nil
}

func (m *mockRepo) ListActiveBySupi(context.Context, string) ([]*domain.AmfSubscription, error) {
	return nil, nil
}

func (m *mockRepo) Update(_ context.Context, sub *domain.AmfSubscription) error {
	m.updates++
	if m.updateErr != nil {
		return m.updateErr
	}
	copied := *sub
	m.sub = &copied
	return nil
}

func (m *mockRepo) Delete(context.Context, string, uint64) error { return nil }

type mockClient struct {
	sent []*model.NamfEventNotification
	err  error
}

func (m *mockClient) SendNotification(_ context.Context, _ string, n *model.NamfEventNotification) error {
	if m.err != nil {
		return m.err
	}
	m.sent = append(m.sent, n)
	return nil
}

func newSubscription(remain *int, expiry *time.Time) *domain.AmfSubscription {
	sub := domain.NewSubscription("a1b2c3d4-0000-4000-8000-000000000001", model.AmfEventSubscription{
		EventList:           []model.AmfEvent{{Type: model.AmfEventTypeLocationReport}},
		EventNotifyUri:      "http://nwdaf:8080/callbacks/amf-notify",
		NotifyCorrelationId: "corr-abc-123",
		NfId:                gppmodel.NfInstanceId(uuid.MustParse("550e8400-e29b-41d4-a716-446655440000")),
		Supi:                "imsi-208950000000010",
	})
	sub.RemainReports = remain
	sub.Expiry = expiry
	return sub
}

func TestDispatchReport_SendsNotification(t *testing.T) {
	remain := 2
	repo := &mockRepo{sub: newSubscription(&remain, nil)}
	client := &mockClient{}
	notifier, err := NewNotifier(repo, client, nil)
	if err != nil {
		t.Fatalf("NewNotifier: %v", err)
	}

	err = notifier.DispatchReport(context.Background(), repo.sub.ID,
		model.NamfEventReport{Type: model.AmfEventTypeLocationReport})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(client.sent) != 1 {
		t.Fatalf("số notification gửi đi = %d, want 1", len(client.sent))
	}
	sent := client.sent[0]
	if sent.NotifyCorrelationId != "corr-abc-123" {
		t.Errorf("notifyCorrelationId = %q", sent.NotifyCorrelationId)
	}
	if len(sent.ReportList) != 1 || sent.ReportList[0].TimeStamp == "" {
		t.Error("report thiếu timeStamp hoặc reportList rỗng")
	}
	if got := sent.ReportList[0].State.RemainReports; got == nil || *got != 1 {
		t.Errorf("remainReports = %v, want 1", got)
	}
}

func TestDispatchReport_NoQuotaLeft(t *testing.T) {
	remain := 0
	repo := &mockRepo{sub: newSubscription(&remain, nil)}
	client := &mockClient{}
	notifier, _ := NewNotifier(repo, client, nil)

	if err := notifier.DispatchReport(context.Background(), repo.sub.ID,
		model.NamfEventReport{Type: model.AmfEventTypeLocationReport}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(client.sent) != 0 {
		t.Error("không được gửi notification khi đã hết quota")
	}
}

func TestDispatchReport_ExpiredSubscription(t *testing.T) {
	past := time.Now().Add(-time.Hour)
	repo := &mockRepo{sub: newSubscription(nil, &past)}
	client := &mockClient{}
	notifier, _ := NewNotifier(repo, client, nil)

	if err := notifier.DispatchReport(context.Background(), repo.sub.ID,
		model.NamfEventReport{Type: model.AmfEventTypeLocationReport}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(client.sent) != 0 {
		t.Error("không được gửi notification cho subscription đã hết hạn")
	}
	if repo.sub.Status != domain.StatusExpired {
		t.Errorf("status = %q, want expired", repo.sub.Status)
	}
}

func TestDispatchReport_UnknownSubscription(t *testing.T) {
	notifier, _ := NewNotifier(&mockRepo{}, &mockClient{}, nil)

	err := notifier.DispatchReport(context.Background(), "a1b2c3d4-0000-4000-8000-000000000009",
		model.NamfEventReport{Type: model.AmfEventTypeLocationReport})
	if !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("err = %v, want ErrNotFound", err)
	}
}

func TestDispatchReport_ClientFailureKeepsQuota(t *testing.T) {
	remain := 3
	repo := &mockRepo{sub: newSubscription(&remain, nil)}
	client := &mockClient{err: errors.New("connection refused")}
	notifier, _ := NewNotifier(repo, client, nil)

	err := notifier.DispatchReport(context.Background(), repo.sub.ID,
		model.NamfEventReport{Type: model.AmfEventTypeLocationReport})
	if err == nil {
		t.Fatal("phải trả lỗi khi consumer không nhận được notification")
	}
	if repo.updates != 0 {
		t.Error("không được ghi lại quota khi gửi thất bại")
	}
}
