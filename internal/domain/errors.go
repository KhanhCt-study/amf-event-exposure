package domain

import (
	"errors"
	"fmt"
	"strings"
)

// Sentinel errors — delivery layer map sang HTTP status:
// ErrInvalid→400, ErrUnsupported→403, ErrNotFound→404, ErrConflict→409, default→500.
var (
	ErrInvalid     = errors.New("invalid request")
	ErrUnsupported = errors.New("unsupported")
	ErrNotFound    = errors.New("not found")
	ErrConflict    = errors.New("conflict")
)

// FieldError gắn một sentinel error với vị trí lỗi trong payload.
// Field luôn là JSON Pointer (RFC 6901), ví dụ "/subscription/eventList".
type FieldError struct {
	Kind   error
	Field  string
	Reason string
}

func NewFieldError(kind error, field, reason string) *FieldError {
	return &FieldError{Kind: kind, Field: field, Reason: reason}
}

func (e *FieldError) Error() string {
	if e.Field == "" {
		return e.Reason
	}
	return fmt.Sprintf("%s: %s", e.Field, e.Reason)
}

func (e *FieldError) Unwrap() error { return e.Kind }

// FieldErrors gom nhiều lỗi validation để build invalidParams trong ProblemDetails.
// Tất cả phần tử phải cùng Kind; Unwrap trả về Kind của phần tử đầu tiên.
type FieldErrors []*FieldError

func (e FieldErrors) Error() string {
	parts := make([]string, 0, len(e))
	for _, fe := range e {
		parts = append(parts, fe.Error())
	}
	return strings.Join(parts, "; ")
}

func (e FieldErrors) Unwrap() error {
	if len(e) == 0 {
		return ErrInvalid
	}
	return e[0].Kind
}

// OrNil trả về nil khi không có lỗi nào — dùng ở cuối hàm validate.
func (e FieldErrors) OrNil() error {
	if len(e) == 0 {
		return nil
	}
	return e
}

// OperationError bọc lỗi kèm ngữ cảnh subscription đang thao tác.
type OperationError struct {
	SubscriptionID string
	Err            error
}

func NewOperationError(subscriptionID string, err error) *OperationError {
	return &OperationError{SubscriptionID: subscriptionID, Err: err}
}

func (e *OperationError) Error() string {
	return fmt.Sprintf("subscription %s: %v", e.SubscriptionID, e.Err)
}

func (e *OperationError) Unwrap() error { return e.Err }
