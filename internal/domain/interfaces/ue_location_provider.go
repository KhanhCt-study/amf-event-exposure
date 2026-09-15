package interfaces

import (
	"context"

	model "github.com/KhanhCt-study/amf-event-exposure.git/pkg/3gppmodel"
)

// IUeLocationProvider trừu tượng hóa nguồn vị trí UE trong AMF.
// Prototype dùng bản static; bản thật sẽ lấy từ UE context / NGAP Location
// Report của AMF.
type IUeLocationProvider interface {
	Location(ctx context.Context, supi string) (*model.UserLocation, error)
}
