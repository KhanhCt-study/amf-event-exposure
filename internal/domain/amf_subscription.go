package domain

import (
	"time"

	model "github.com/KhanhCt-study/amf-event-exposure.git/pkg/3gppmodel"
)

// SubscriptionStatus — vòng đời của một subscription phía AMF.
type SubscriptionStatus string

const (
	StatusActive    SubscriptionStatus = "active"
	StatusExpired   SubscriptionStatus = "expired"
	StatusCompleted SubscriptionStatus = "completed" // đã gửi đủ maxReports
)

// AmfSubscription là entity lưu trong bảng amf_subscriptions.
// Data giữ nguyên AmfEventSubscription (JSONB); các field còn lại là cột
// denormalized phục vụ truy vấn và điều kiện CAS.
type AmfSubscription struct {
	ID      string
	Version uint64
	Status  SubscriptionStatus

	Data model.AmfEventSubscription

	EventNotifyUri      string
	NotifyCorrelationId string
	NfId                string
	Supi                *string

	TriggerType   model.AmfEventTrigger
	MaxReports    *int
	RemainReports *int
	Expiry        *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewSubscription dựng entity từ payload đã validate. Version bắt đầu từ 1
// (schema có CHECK version > 0).
func NewSubscription(id string, data model.AmfEventSubscription) *AmfSubscription {
	sub := &AmfSubscription{
		ID:      id,
		Version: 1,
		Status:  StatusActive,
		Data:    data,
	}
	sub.SyncFromData()
	return sub
}

// SyncFromData đồng bộ các cột denormalized theo Data. Bắt buộc gọi sau mỗi
// lần sửa Data, nếu không CHECK constraint trong DB sẽ từ chối row.
func (s *AmfSubscription) SyncFromData() {
	s.EventNotifyUri = s.Data.EventNotifyUri
	s.NotifyCorrelationId = s.Data.NotifyCorrelationId
	s.NfId = s.Data.NfId.String()

	if s.Data.Supi != "" {
		supi := s.Data.Supi
		s.Supi = &supi
	} else {
		s.Supi = nil
	}

	s.TriggerType = model.AmfEventTriggerContinuous
	s.MaxReports = nil
	s.Expiry = nil

	if opt := s.Data.Options; opt != nil {
		if opt.Trigger != "" {
			s.TriggerType = opt.Trigger
		}
		if opt.MaxReports > 0 {
			max := opt.MaxReports
			s.MaxReports = &max
		}
		if opt.Expiry != nil {
			expiry := *opt.Expiry
			s.Expiry = &expiry
		}
	}

	if s.MaxReports == nil {
		s.RemainReports = nil
		return
	}
	// Giữ nguyên RemainReports nếu đã đếm dở và vẫn còn hợp lệ.
	if s.RemainReports == nil || *s.RemainReports > *s.MaxReports {
		remain := s.MaxReports
		s.RemainReports = remain
	}
}

// IsExpired kiểm tra options.expiry đã qua hay chưa.
func (s *AmfSubscription) IsExpired(now time.Time) bool {
	return s.Expiry != nil && !now.Before(*s.Expiry)
}

// ConsumeReport trừ một lượt report. Trả về false khi subscription đã hết
// quota — caller không được gửi notification nữa.
func (s *AmfSubscription) ConsumeReport() bool {
	if s.Status != StatusActive {
		return false
	}
	if s.RemainReports == nil {
		return true
	}
	if *s.RemainReports <= 0 {
		s.Status = StatusCompleted
		return false
	}
	*s.RemainReports--
	if s.RemainReports == nil {
		s.Status = StatusCompleted
	}
	return true
}

// HasEventType cho biết subscription có đăng ký loại event này không.
func (s *AmfSubscription) HasEventType(t model.AmfEventType) bool {
	for i := range s.Data.EventList {
		if s.Data.EventList[i].Type == t {
			return true
		}
	}
	return false
}
