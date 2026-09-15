package interfaces

import (
	"context"

	model "github.com/KhanhCt-study/amf-event-exposure.git/pkg/3gppmodel"
)

// IAmfEventSubUseCase — business logic của Namf_EventExposure (TS 29.518 §5.3).
type IAmfEventSubUseCase interface {
	CreateSubscription(ctx context.Context, req *model.AmfCreateEventSubscription) (*model.AmfCreatedEventSubscription, error)
	ModifySubscription(ctx context.Context, id string, patches []model.AmfUpdateEventSubscriptionItem) (*model.AmfUpdatedEventSubscription, error)
	DeleteSubscription(ctx context.Context, id string) error
}

// INotificationUseCase — dispatch event report đến consumer đã đăng ký.
type INotificationUseCase interface {
	DispatchReport(ctx context.Context, subscriptionID string, report model.NamfEventReport) error
}
