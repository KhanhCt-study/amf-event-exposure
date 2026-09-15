// Package notification dựng và gửi NamfEventNotification đến NF consumer.
package notification

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/KhanhCt-study/amf-event-exposure.git/internal/domain"
	"github.com/KhanhCt-study/amf-event-exposure.git/internal/domain/interfaces"
	model "github.com/KhanhCt-study/amf-event-exposure.git/pkg/3gppmodel"
)

var _ interfaces.INotificationUseCase = (*NotifierImpl)(nil)

// NotifierImpl gửi event report cho đúng subscription và cập nhật quota.
type NotifierImpl struct {
	repo   interfaces.IAmfSubRepository
	client interfaces.INotificationClient
	logger *slog.Logger
	now    func() time.Time
}

// NewNotifier dựng notifier với constructor injection.
func NewNotifier(
	repo interfaces.IAmfSubRepository,
	client interfaces.INotificationClient,
	logger *slog.Logger,
) (*NotifierImpl, error) {
	if repo == nil {
		return nil, errors.New("notification: repository must not be nil")
	}
	if client == nil {
		return nil, errors.New("notification: client must not be nil")
	}
	if logger == nil {
		logger = slog.Default()
	}
	return &NotifierImpl{repo: repo, client: client, logger: logger, now: time.Now}, nil
}

// DispatchReport gửi một report đến consumer của subscription chỉ định.
// Luồng: load subscription → kiểm tra trạng thái/quota → POST callback →
// ghi lại remainReports và status.
func (n *NotifierImpl) DispatchReport(
	ctx context.Context,
	subscriptionID string,
	report model.NamfEventReport,
) error {
	sub, err := n.repo.GetByID(ctx, subscriptionID)
	if err != nil {
		return domain.NewOperationError(subscriptionID, err)
	}

	if !sub.HasEventType(report.Type) {
		return domain.NewOperationError(subscriptionID,
			domain.NewFieldError(domain.ErrInvalid, "/type", "event type is not subscribed"))
	}

	now := n.now().UTC()
	if sub.IsExpired(now) {
		sub.Status = domain.StatusExpired
		sub.UpdatedAt = now
		if err := n.repo.Update(ctx, sub); err != nil {
			n.logger.Warn("cannot mark subscription expired", "subscriptionId", sub.ID, "error", err)
		}
		return nil
	}

	if !sub.ConsumeReport() {
		n.logger.Debug("subscription has no report quota left", "subscriptionId", sub.ID)
		return nil
	}

	if report.TimeStamp == "" {
		report.TimeStamp = now.String()
	}
	report.State = model.NamfEventState{
		Active:        sub.Status == domain.StatusActive,
		RemainReports: sub.RemainReports,
	}

	notification := &model.NamfEventNotification{
		NotifyCorrelationId: sub.NotifyCorrelationId,
		ReportList:          []model.NamfEventReport{report},
	}

	if err := n.client.SendNotification(ctx, sub.EventNotifyUri, notification); err != nil {
		// Không trừ quota khi gửi thất bại.
		return domain.NewOperationError(sub.ID, err)
	}

	sub.UpdatedAt = now
	if err := n.repo.Update(ctx, sub); err != nil {
		n.logger.Warn("notification sent but state not persisted",
			"subscriptionId", sub.ID, "error", err)
	}
	return nil
}
