// Package subscription chứa business logic của AMF Event Exposure:
// validate payload, lưu trữ subscription và dựng response theo TS 29.518.
package subscription

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/KhanhCt-study/amf-event-exposure.git/internal/domain"
	"github.com/KhanhCt-study/amf-event-exposure.git/internal/domain/interfaces"
	model "github.com/KhanhCt-study/amf-event-exposure.git/pkg/3gppmodel"
)

// SubscriptionsPath là base path của resource collection (TS 29.518 §6.2.3.2).
const SubscriptionsPath = "/namf-evts/v1/subscriptions"

var _ interfaces.IAmfEventSubUseCase = (*AmfSubUseCaseImpl)(nil)

// AmfSubUseCaseImpl triển khai IAmfEventSubUseCase.
type AmfSubUseCaseImpl struct {
	repo         interfaces.IAmfSubRepository
	locations    interfaces.IUeLocationProvider
	publicAPIURL string
	now          func() time.Time
	newID        func() string
}

// NewAmfSubUseCase dựng usecase với constructor injection.
// publicAPIURL là URL public của AMF (dùng cho Location header và
// subscriptionId dạng URI đầy đủ).
func NewAmfSubUseCase(
	repo interfaces.IAmfSubRepository,
	locations interfaces.IUeLocationProvider,
	publicAPIURL string,
) (*AmfSubUseCaseImpl, error) {
	if repo == nil {
		return nil, errors.New("subscription: repository must not be nil")
	}
	if locations == nil {
		return nil, errors.New("subscription: location provider must not be nil")
	}
	if strings.TrimSpace(publicAPIURL) == "" {
		return nil, errors.New("subscription: publicAPIURL must not be empty")
	}
	return &AmfSubUseCaseImpl{
		repo:         repo,
		locations:    locations,
		publicAPIURL: strings.TrimRight(publicAPIURL, "/"),
		now:          time.Now,
		newID:        func() string { return uuid.NewString() },
	}, nil
}

// ResourceURI trả về URI đầy đủ của một subscription.
func (uc *AmfSubUseCaseImpl) ResourceURI(id string) string {
	return uc.publicAPIURL + SubscriptionsPath + "/" + id
}

// CreateSubscription — POST /namf-evts/v1/subscriptions.
func (uc *AmfSubUseCaseImpl) CreateSubscription(ctx context.Context, req *model.AmfCreateEventSubscription) (*model.AmfCreatedEventSubscription, error) {
	now := uc.now().UTC()
	if err := ValidateCreate(req, now); err != nil {
		return nil, err
	}

	sub := domain.NewSubscription(uc.newID(), req.Subscription)
	sub.CreatedAt, sub.UpdatedAt = now, now

	if err := uc.repo.Create(ctx, sub); err != nil {
		return nil, domain.NewOperationError(sub.ID, err)
	}

	resp := &model.AmfCreatedEventSubscription{
		Subscription:      sub.Data,
		SubscriptionId:    uc.ResourceURI(sub.ID),
		SupportedFeatures: req.SupportedFeatures,
	}
	if reports := uc.immediateReports(ctx, sub, now); len(reports) > 0 {
		resp.ReportList = reports
	}
	return resp, nil
}

// ModifySubscription — PATCH /namf-evts/v1/subscriptions/{subscriptionId}.
func (uc *AmfSubUseCaseImpl) ModifySubscription(
	ctx context.Context,
	id string,
	patches []model.AmfUpdateEventSubscriptionItem,
) (*model.AmfUpdatedEventSubscription, error) {
	if err := validateID(id); err != nil {
		return nil, err
	}
	if err := ValidatePatches(patches); err != nil {
		return nil, err
	}

	sub, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, domain.NewOperationError(id, err)
	}

	updated := sub.Data
	for i := range patches {
		if err := applyPatch(&updated, &patches[i], i); err != nil {
			return nil, domain.NewOperationError(id, err)
		}
	}

	now := uc.now().UTC()
	if err := ValidateCreate(&model.AmfCreateEventSubscription{Subscription: updated}, now); err != nil {
		return nil, domain.NewOperationError(id, err)
	}

	sub.Data = updated
	sub.UpdatedAt = now
	sub.SyncFromData()

	if err := uc.repo.Update(ctx, sub); err != nil {
		return nil, domain.NewOperationError(id, err)
	}

	return &model.AmfUpdatedEventSubscription{
		Subscription: sub.Data,
		ReportList:   uc.immediateReports(ctx, sub, now),
	}, nil
}

// DeleteSubscription — DELETE /namf-evts/v1/subscriptions/{subscriptionId}.
func (uc *AmfSubUseCaseImpl) DeleteSubscription(ctx context.Context, id string) error {
	if err := validateID(id); err != nil {
		return err
	}
	sub, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return domain.NewOperationError(id, err)
	}
	if err := uc.repo.Delete(ctx, id, sub.Version); err != nil {
		return domain.NewOperationError(id, err)
	}
	return nil
}

// immediateReports dựng report trả ngay cho các event có immediateFlag=true
// (TS 29.518 §5.3.2.2.2).
func (uc *AmfSubUseCaseImpl) immediateReports(
	ctx context.Context,
	sub *domain.AmfSubscription,
	now time.Time,
) []model.NamfEventReport {
	var reports []model.NamfEventReport
	for i := range sub.Data.EventList {
		ev := &sub.Data.EventList[i]
		if !ev.ImmediateFlag || ev.Type != model.AmfEventTypeLocationReport {
			continue
		}
		location, err := uc.locations.Location(ctx, sub.Data.Supi)
		if err != nil {
			// Không có vị trí thì bỏ qua immediate report, subscription vẫn hợp lệ.
			continue
		}
		reports = append(reports, model.NamfEventReport{
			Type:      model.AmfEventTypeLocationReport,
			State:     model.NamfEventState{Active: sub.Status == domain.StatusActive, RemainReports: sub.RemainReports},
			TimeStamp: now.String(),
			Supi:      sub.Data.Supi,
			AnyUe:     &sub.Data.AnyUE,
			Location:  location,
		})
	}
	return reports
}

func validateID(id string) error {
	if id == "" {
		return domain.NewFieldError(domain.ErrInvalid, "/subscriptionId", "subscriptionId is required")
	}
	if _, err := uuid.Parse(id); err != nil {
		return domain.NewFieldError(domain.ErrInvalid, "/subscriptionId", "must be a valid UUID")
	}
	return nil
}

// applyPatch áp dụng một thao tác RFC 6902 lên AmfEventSubscription.
// Chỉ các path trong phạm vi prototype được chấp nhận; path khác → 403.
func applyPatch(sub *model.AmfEventSubscription, patch *model.AmfUpdateEventSubscriptionItem, index int) error {
	ptr := fmt.Sprintf("/%d/path", index)
	path := strings.TrimPrefix(patch.Path, "/subscription")

	switch {
	case path == "/options":
		if patch.Op != "replace" {
			return domain.NewFieldError(domain.ErrUnsupported, ptr, "only replace is supported on /subscription/options")
		}
		var options model.AmfEventMode
		if err := decodeValue(patch.Value, &options, index); err != nil {
			return err
		}
		sub.Options = &options
		return nil

	case path == "/eventList":
		if patch.Op != "replace" {
			return domain.NewFieldError(domain.ErrUnsupported, ptr, "only replace is supported on /subscription/eventList")
		}
		var events []model.AmfEvent
		if err := decodeValue(patch.Value, &events, index); err != nil {
			return err
		}
		sub.EventList = events
		return nil

	case strings.HasPrefix(path, "/eventList/"):
		return applyEventListPatch(sub, patch, index, strings.TrimPrefix(path, "/eventList/"))

	default:
		return domain.NewFieldError(domain.ErrUnsupported, ptr,
			"prototype supports only /subscription/eventList and /subscription/options")
	}
}

func applyEventListPatch(
	sub *model.AmfEventSubscription,
	patch *model.AmfUpdateEventSubscriptionItem,
	index int,
	token string,
) error {
	ptr := fmt.Sprintf("/%d/path", index)

	if token == "-" {
		if patch.Op != "add" {
			return domain.NewFieldError(domain.ErrInvalid, ptr, "'-' index is only valid with op add")
		}
		var event model.AmfEvent
		if err := decodeValue(patch.Value, &event, index); err != nil {
			return err
		}
		sub.EventList = append(sub.EventList, event)
		return nil
	}

	i, err := strconv.Atoi(token)
	if err != nil || i < 0 {
		return domain.NewFieldError(domain.ErrInvalid, ptr, "eventList index must be a non-negative integer or '-'")
	}

	switch patch.Op {
	case "add":
		if i > len(sub.EventList) {
			return domain.NewFieldError(domain.ErrInvalid, ptr, "index out of range")
		}
		var event model.AmfEvent
		if err := decodeValue(patch.Value, &event, index); err != nil {
			return err
		}
		sub.EventList = append(sub.EventList, model.AmfEvent{})
		copy(sub.EventList[i+1:], sub.EventList[i:])
		sub.EventList[i] = event
		return nil

	case "replace":
		if i >= len(sub.EventList) {
			return domain.NewFieldError(domain.ErrInvalid, ptr, "index out of range")
		}
		var event model.AmfEvent
		if err := decodeValue(patch.Value, &event, index); err != nil {
			return err
		}
		sub.EventList[i] = event
		return nil

	case "remove":
		if i >= len(sub.EventList) {
			return domain.NewFieldError(domain.ErrInvalid, ptr, "index out of range")
		}
		sub.EventList = append(sub.EventList[:i], sub.EventList[i+1:]...)
		return nil
	}
	return domain.NewFieldError(domain.ErrInvalid, fmt.Sprintf("/%d/op", index), "unsupported operation")
}

// decodeValue chuyển patch.Value (interface{}) sang kiểu 3GPP tương ứng,
// vẫn giữ nguyên tắc strict: field lạ bị từ chối.
func decodeValue(value interface{}, dst interface{}, index int) error {
	ptr := fmt.Sprintf("/%d/value", index)
	if value == nil {
		return domain.NewFieldError(domain.ErrInvalid, ptr, "value is required")
	}
	raw, err := json.Marshal(value)
	if err != nil {
		return domain.NewFieldError(domain.ErrInvalid, ptr, "cannot encode patch value")
	}
	dec := json.NewDecoder(strings.NewReader(string(raw)))
	dec.DisallowUnknownFields()
	if err := dec.Decode(dst); err != nil {
		return domain.NewFieldError(domain.ErrInvalid, ptr, "patch value does not match the target schema")
	}
	return nil
}
