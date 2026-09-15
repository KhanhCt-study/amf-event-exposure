// Package client chứa các HTTP client gọi ra NF khác.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/KhanhCt-study/amf-event-exposure.git/internal/domain"
	"github.com/KhanhCt-study/amf-event-exposure.git/internal/domain/interfaces"
	model "github.com/KhanhCt-study/amf-event-exposure.git/pkg/3gppmodel"
)

var _ interfaces.INotificationClient = (*NotificationHTTPClient)(nil)

// maxErrorBody giới hạn phần body lỗi đọc về để log.
const maxErrorBody = 4 << 10

// NotificationHTTPClient POST NamfEventNotification đến eventNotifyUri.
type NotificationHTTPClient struct {
	httpClient *http.Client
}

// NewNotificationHTTPClient dựng client với timeout chỉ định.
func NewNotificationHTTPClient(timeout time.Duration) (*NotificationHTTPClient, error) {
	if timeout <= 0 {
		return nil, errors.New("client: timeout must be greater than 0")
	}
	return &NotificationHTTPClient{httpClient: &http.Client{Timeout: timeout}}, nil
}

// SendNotification gửi callback; consumer trả 2xx (thường là 204) là thành công.
func (c *NotificationHTTPClient) SendNotification(
	ctx context.Context,
	notifyURI string,
	notification *model.NamfEventNotification,
) error {
	if notifyURI == "" {
		return fmt.Errorf("%w: eventNotifyUri is empty", domain.ErrInvalid)
	}
	if notification == nil {
		return fmt.Errorf("%w: notification must not be nil", domain.ErrInvalid)
	}

	body, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("marshal notification: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, notifyURI, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("build notification request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/problem+json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("send notification to %s: %w", notifyURI, err)
	}
	defer func() {
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	detail, _ := io.ReadAll(io.LimitReader(resp.Body, maxErrorBody))
	return fmt.Errorf("consumer %s rejected notification: status %d: %s", notifyURI, resp.StatusCode, string(detail))
}
