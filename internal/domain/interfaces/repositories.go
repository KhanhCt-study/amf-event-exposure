package interfaces

import (
	"context"

	"github.com/KhanhCt-study/amf-event-exposure.git/internal/domain"
)

// IAmfSubReader — phần read-only, dùng cho notification pipeline.
type IAmfSubReader interface {
	GetByID(ctx context.Context, id string) (*domain.AmfSubscription, error)
	ListActiveBySupi(ctx context.Context, supi string) ([]*domain.AmfSubscription, error)
}

// IAmfSubRepository — CRUD trên bảng amf_subscriptions.
// Update/Delete dùng CAS: so khớp sub.Version (Update) hoặc expectedVersion
// (Delete) với version đang có trong DB; lệch → domain.ErrConflict.
type IAmfSubRepository interface {
	IAmfSubReader

	Create(ctx context.Context, sub *domain.AmfSubscription) error
	// Update tăng version trong DB và ghi version mới ngược lại vào sub.
	Update(ctx context.Context, sub *domain.AmfSubscription) error
	Delete(ctx context.Context, id string, expectedVersion uint64) error
}
