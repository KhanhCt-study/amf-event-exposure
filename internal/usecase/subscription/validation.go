package subscription

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/KhanhCt-study/amf-event-exposure.git/internal/delivery/http/validator"
	"github.com/KhanhCt-study/amf-event-exposure.git/internal/domain"
	gppmodel "github.com/KhanhCt-study/amf-event-exposure.git/pkg/3gppmodel"
	model "github.com/KhanhCt-study/amf-event-exposure.git/pkg/3gppmodel"
	"github.com/google/uuid"
)

// Các field được prototype hỗ trợ. Field nằm ngoài danh sách này tuy hợp lệ
// theo TS 29.518 nhưng chưa triển khai → 403 Forbidden.
var (
	supportedSubscriptionFields = map[string]bool{
		"eventList":           true,
		"eventNotifyUri":      true,
		"notifyCorrelationId": true,
		"nfId":                true,
		"supi":                true,
		"anyUE":               true,
		"options":             true,
		"sourceNfType":        true,
	}
	supportedEventFields = map[string]bool{
		"type":            true,
		"immediateFlag":   true,
		"maxReports":      true,
		"maxResponseTime": true,
	}
	supportedTriggers = map[model.AmfEventTrigger]bool{
		model.AmfEventTriggerOneTime:    true,
		model.AmfEventTriggerContinuous: true,
		model.AmfEventTriggerPeriodic:   true,
	}
)

// ValidateCreate kiểm tra AmfCreateEventSubscription theo TS 29.518 §6.2.6.2.2
// và giới hạn của prototype.
func ValidateCreate(req *model.AmfCreateEventSubscription, now time.Time) error {
	if req == nil {
		return domain.NewFieldError(domain.ErrInvalid, "/subscription", "request body is required")
	}
	sub := &req.Subscription

	// Field chưa hỗ trợ → 403 (fatal, trả ngay).
	if err := rejectUnsupportedFields(sub); err != nil {
		return err
	}

	var errs domain.FieldErrors

	if len(sub.EventList) == 0 {
		errs = append(errs, domain.NewFieldError(domain.ErrInvalid,
			"/subscription/eventList", "at least one event is required"))
	}
	for i := range sub.EventList {
		ptr := fmt.Sprintf("/subscription/eventList/%d", i)
		ev := &sub.EventList[i]

		switch {
		case ev.Type == "":
			errs = append(errs, domain.NewFieldError(domain.ErrInvalid, ptr+"/type", "event type is required"))
		case ev.Type != model.AmfEventTypeLocationReport:
			return domain.NewFieldError(domain.ErrUnsupported, ptr+"/type",
				fmt.Sprintf("prototype supports only %s", model.AmfEventTypeLocationReport))
		}
		if ev.MaxReports != 0 && ev.MaxReports <= 0 {
			errs = append(errs, domain.NewFieldError(domain.ErrInvalid, ptr+"/maxReports", "must be greater than 0"))
		}
		if ev.MaxResponseTime != 0 && ev.MaxResponseTime < 0 {
			errs = append(errs, domain.NewFieldError(domain.ErrInvalid, ptr+"/maxResponseTime", "must not be negative"))
		}
	}

	if sub.EventNotifyUri == "" {
		errs = append(errs, domain.NewFieldError(domain.ErrInvalid,
			"/subscription/eventNotifyUri", "eventNotifyUri is required"))
	} else if !validator.IsAbsoluteHTTPURL(sub.EventNotifyUri) {
		errs = append(errs, domain.NewFieldError(domain.ErrInvalid,
			"/subscription/eventNotifyUri", "must be an absolute http(s) URL"))
	}

	if strings.TrimSpace(sub.NotifyCorrelationId) == "" {
		errs = append(errs, domain.NewFieldError(domain.ErrInvalid,
			"/subscription/notifyCorrelationId", "notifyCorrelationId is required"))
	}

	// Sửa dòng if cũ thành kiểm tra trực tiếp với uuid.Nil
	if sub.NfId == gppmodel.NfInstanceId(uuid.Nil) {
		errs = append(errs, domain.NewFieldError(domain.ErrInvalid, "/subscription/nfId", "nfId is required"))
	} else if !validator.IsUUID(sub.NfId.String()) {
		errs = append(errs, domain.NewFieldError(domain.ErrInvalid, "/subscription/nfId", "must be a valid UUID"))
	}

	if sub.Supi != "" && !validator.IsSupi(sub.Supi) {
		errs = append(errs, domain.NewFieldError(domain.ErrInvalid,
			"/subscription/supi", "must match ^imsi-[0-9]{5,15}$"))
	}
	if sub.Supi == "" && !sub.AnyUE {
		errs = append(errs, domain.NewFieldError(domain.ErrInvalid,
			"/subscription/supi", "either supi or anyUE is required"))
	}

	errs = append(errs, validateOptions(sub.Options, now)...)
	return errs.OrNil()
}

func validateOptions(opt *model.AmfEventMode, now time.Time) domain.FieldErrors {
	if opt == nil {
		return nil
	}
	var errs domain.FieldErrors

	if opt.Trigger == "" {
		errs = append(errs, domain.NewFieldError(domain.ErrInvalid,
			"/subscription/options/trigger", "trigger is required when options is present"))
	} else if !supportedTriggers[opt.Trigger] {
		errs = append(errs, domain.NewFieldError(domain.ErrInvalid, "/subscription/options/trigger",
			"must be one of ONE_TIME, CONTINUOUS, PERIODIC"))
	}
	if opt.MaxReports != 0 && opt.MaxReports <= 0 {
		errs = append(errs, domain.NewFieldError(domain.ErrInvalid,
			"/subscription/options/maxReports", "must be greater than 0"))
	}
	if opt.Expiry != nil && !opt.Expiry.After(now) {
		errs = append(errs, domain.NewFieldError(domain.ErrInvalid,
			"/subscription/options/expiry", "must be in the future"))
	}
	if opt.Trigger == model.AmfEventTriggerPeriodic && (opt.RepPeriod == 0 || opt.RepPeriod <= 0) {
		errs = append(errs, domain.NewFieldError(domain.ErrInvalid,
			"/subscription/options/repPeriod", "required and must be greater than 0 for PERIODIC trigger"))
	}
	return errs
}

// rejectUnsupportedFields dùng kỹ thuật JSON round-trip: marshal lại struct đã
// decode rồi đối chiếu key với whitelist. Cách này bắt được các field hợp lệ
// theo spec nhưng prototype chưa xử lý (giống validation.go bên NWDAF).
func rejectUnsupportedFields(sub *model.AmfEventSubscription) error {
	raw, err := json.Marshal(sub)
	if err != nil {
		return domain.NewFieldError(domain.ErrInvalid, "/subscription", "cannot re-encode subscription")
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(raw, &fields); err != nil {
		return domain.NewFieldError(domain.ErrInvalid, "/subscription", "cannot inspect subscription fields")
	}
	for key := range fields {
		if !supportedSubscriptionFields[key] {
			return domain.NewFieldError(domain.ErrUnsupported,
				"/subscription/"+validator.EscapePointerSegment(key), "field is not supported by this prototype")
		}
	}

	for i := range sub.EventList {
		evRaw, err := json.Marshal(&sub.EventList[i])
		if err != nil {
			return domain.NewFieldError(domain.ErrInvalid,
				fmt.Sprintf("/subscription/eventList/%d", i), "cannot re-encode event")
		}
		var evFields map[string]json.RawMessage
		if err := json.Unmarshal(evRaw, &evFields); err != nil {
			return domain.NewFieldError(domain.ErrInvalid,
				fmt.Sprintf("/subscription/eventList/%d", i), "cannot inspect event fields")
		}
		for key := range evFields {
			if !supportedEventFields[key] {
				return domain.NewFieldError(domain.ErrUnsupported,
					fmt.Sprintf("/subscription/eventList/%d/%s", i, validator.EscapePointerSegment(key)),
					"field is not supported by this prototype")
			}
		}
	}
	return nil
}

// ValidatePatches kiểm tra danh sách thao tác PATCH (RFC 6902).
// Prototype chỉ hỗ trợ:
//   - replace /subscription/eventList/{index}
//   - add     /subscription/eventList/-
//   - remove  /subscription/eventList/{index}
//   - replace /subscription/options
func ValidatePatches(patches []model.AmfUpdateEventSubscriptionItem) error {
	if len(patches) == 0 {
		return domain.NewFieldError(domain.ErrInvalid, "", "at least one patch operation is required")
	}

	var errs domain.FieldErrors
	for i := range patches {
		ptr := fmt.Sprintf("/%d", i)
		op := patches[i].Op

		switch op {
		case "add", "remove", "replace":
		case "":
			errs = append(errs, domain.NewFieldError(domain.ErrInvalid, ptr+"/op", "op is required"))
			continue
		default:
			errs = append(errs, domain.NewFieldError(domain.ErrInvalid, ptr+"/op",
				"must be one of add, remove, replace"))
			continue
		}

		path := patches[i].Path
		if path == "" {
			errs = append(errs, domain.NewFieldError(domain.ErrInvalid, ptr+"/path", "path is required"))
			continue
		}
		if !validator.IsJSONPointer(path) {
			errs = append(errs, domain.NewFieldError(domain.ErrInvalid, ptr+"/path",
				"must be a valid JSON Pointer (RFC 6901)"))
			continue
		}
		if op != "remove" && patches[i].Value == nil {
			errs = append(errs, domain.NewFieldError(domain.ErrInvalid, ptr+"/value",
				"value is required for add and replace"))
		}
	}
	return errs.OrNil()
}
