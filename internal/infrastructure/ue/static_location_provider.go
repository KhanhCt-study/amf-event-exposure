// Package ue mô phỏng nguồn dữ liệu UE của AMF trong giai đoạn prototype.
package ue

import (
	"context"

	"github.com/KhanhCt-study/amf-event-exposure.git/internal/domain/interfaces"
	model "github.com/KhanhCt-study/amf-event-exposure.git/pkg/3gppmodel"
)

var _ interfaces.IUeLocationProvider = (*StaticLocationProvider)(nil)

// StaticLocationProvider trả về một NR location cố định lấy từ config.
// Khi AMF có UE context thật, thay implementation này mà không đụng usecase.
type StaticLocationProvider struct {
	mcc      string
	mnc      string
	tac      string
	nrCellId string
}

// NewStaticLocationProvider dựng provider từ PLMN và cell mặc định.
func NewStaticLocationProvider(mcc, mnc, tac, nrCellId string) *StaticLocationProvider {
	return &StaticLocationProvider{mcc: mcc, mnc: mnc, tac: tac, nrCellId: nrCellId}
}

// Location trả về vị trí hiện tại của UE (prototype: giá trị tĩnh).
func (p *StaticLocationProvider) Location(_ context.Context, _ string) (*model.UserLocation, error) {
	plmn := model.PlmnId{Mcc: p.mcc, Mnc: p.mnc}
	return &model.UserLocation{
		NrLocation: &model.NrLocation{
			Tai:  model.Tai{PlmnId: plmn, Tac: p.tac},
			Ncgi: model.Ncgi{PlmnId: plmn, NrCellId: p.nrCellId},
		},
	}, nil
}
