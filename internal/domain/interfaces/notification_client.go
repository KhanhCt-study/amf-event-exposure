package interfaces

import (
	"context"

	model "github.com/KhanhCt-study/amf-event-exposure.git/pkg/3gppmodel"
)

// INotificationClient gửi POST {eventNotifyUri} đến NF consumer.
// Consumer trả 204 No Content khi nhận thành công.
type INotificationClient interface {
	SendNotification(ctx context.Context, notifyURI string, notification *model.NamfEventNotification) error
}
