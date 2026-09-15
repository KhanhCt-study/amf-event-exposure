// Package database triển khai repository trên PostgreSQL bằng pgx (không ORM).
package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/KhanhCt-study/amf-event-exposure.git/internal/domain"
	"github.com/KhanhCt-study/amf-event-exposure.git/internal/domain/interfaces"
	model "github.com/KhanhCt-study/amf-event-exposure.git/pkg/3gppmodel"
)

var _ interfaces.IAmfSubRepository = (*PostgresAmfSubRepo)(nil)

const selectColumns = `
	subscription_id, version, status, subscription_data,
	event_notify_uri, notify_correlation_id, nf_id, supi,
	trigger_type, max_reports, remain_reports, expiry,
	created_at, updated_at`

// PostgresAmfSubRepo lưu subscription trong bảng amf_subscriptions.
type PostgresAmfSubRepo struct {
	pool *pgxpool.Pool
}

// NewPostgresAmfSubRepo dựng repository từ connection pool đã khởi tạo.
func NewPostgresAmfSubRepo(pool *pgxpool.Pool) (*PostgresAmfSubRepo, error) {
	if pool == nil {
		return nil, errors.New("database: pool must not be nil")
	}
	return &PostgresAmfSubRepo{pool: pool}, nil
}

// Create chèn subscription mới (version = 1).
func (r *PostgresAmfSubRepo) Create(ctx context.Context, sub *domain.AmfSubscription) error {
	if sub == nil {
		return fmt.Errorf("%w: subscription must not be nil", domain.ErrInvalid)
	}
	data, err := json.Marshal(sub.Data)
	if err != nil {
		return fmt.Errorf("marshal subscription_data: %w", err)
	}

	const query = `
		INSERT INTO amf_subscriptions (
			subscription_id, version, status, subscription_data,
			event_notify_uri, notify_correlation_id, nf_id, supi,
			trigger_type, max_reports, remain_reports, expiry
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING created_at, updated_at`

	err = r.pool.QueryRow(ctx, query,
		sub.ID, int64(sub.Version), string(sub.Status), data,
		sub.EventNotifyUri, sub.NotifyCorrelationId, sub.NfId, sub.Supi,
		string(sub.TriggerType), sub.MaxReports, sub.RemainReports, sub.Expiry,
	).Scan(&sub.CreatedAt, &sub.UpdatedAt)
	if err != nil {
		return fmt.Errorf("insert subscription: %w", err)
	}
	return nil
}

// GetByID trả về subscription theo id; không tồn tại → domain.ErrNotFound.
func (r *PostgresAmfSubRepo) GetByID(ctx context.Context, id string) (*domain.AmfSubscription, error) {
	query := `SELECT` + selectColumns + ` FROM amf_subscriptions WHERE subscription_id = $1`

	sub, err := scanSubscription(r.pool.QueryRow(ctx, query, id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%w: subscription %s", domain.ErrNotFound, id)
		}
		return nil, fmt.Errorf("select subscription: %w", err)
	}
	return sub, nil
}

// ListActiveBySupi trả về các subscription đang active của một SUPI.
func (r *PostgresAmfSubRepo) ListActiveBySupi(ctx context.Context, supi string) ([]*domain.AmfSubscription, error) {
	query := `SELECT` + selectColumns + `
		FROM amf_subscriptions
		WHERE status = 'active' AND (supi = $1 OR supi IS NULL)`

	rows, err := r.pool.Query(ctx, query, supi)
	if err != nil {
		return nil, fmt.Errorf("select subscriptions by supi: %w", err)
	}
	defer rows.Close()

	var subs []*domain.AmfSubscription
	for rows.Next() {
		sub, err := scanSubscription(rows)
		if err != nil {
			return nil, fmt.Errorf("scan subscription: %w", err)
		}
		subs = append(subs, sub)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate subscriptions: %w", err)
	}
	return subs, nil
}

// Update ghi đè subscription với optimistic concurrency: chỉ thành công khi
// version trong DB bằng sub.Version. Version mới được ghi ngược vào sub.
func (r *PostgresAmfSubRepo) Update(ctx context.Context, sub *domain.AmfSubscription) error {
	if sub == nil {
		return fmt.Errorf("%w: subscription must not be nil", domain.ErrInvalid)
	}
	data, err := json.Marshal(sub.Data)
	if err != nil {
		return fmt.Errorf("marshal subscription_data: %w", err)
	}

	const query = `
		WITH target AS (
			SELECT subscription_id, version
			FROM amf_subscriptions
			WHERE subscription_id = $1
		), updated AS (
			UPDATE amf_subscriptions s
			SET version               = s.version + 1,
			    status                = $3,
			    subscription_data     = $4,
			    event_notify_uri      = $5,
			    notify_correlation_id = $6,
			    nf_id                 = $7,
			    supi                  = $8,
			    trigger_type          = $9,
			    max_reports           = $10,
			    remain_reports        = $11,
			    expiry                = $12,
			    updated_at            = clock_timestamp()
			FROM target t
			WHERE s.subscription_id = t.subscription_id AND t.version = $2
			RETURNING s.version, s.updated_at
		)
		SELECT
			(SELECT count(*) FROM target)  AS found,
			(SELECT version    FROM updated) AS new_version,
			(SELECT updated_at FROM updated) AS updated_at`

	var (
		found      int64
		newVersion *int64
		updatedAt  *time.Time
	)
	err = r.pool.QueryRow(ctx, query,
		sub.ID, int64(sub.Version), string(sub.Status), data,
		sub.EventNotifyUri, sub.NotifyCorrelationId, sub.NfId, sub.Supi,
		string(sub.TriggerType), sub.MaxReports, sub.RemainReports, sub.Expiry,
	).Scan(&found, &newVersion, &updatedAt)
	if err != nil {
		return fmt.Errorf("update subscription: %w", err)
	}

	switch {
	case found == 0:
		return fmt.Errorf("%w: subscription %s", domain.ErrNotFound, sub.ID)
	case newVersion == nil:
		return fmt.Errorf("%w: subscription %s was modified concurrently", domain.ErrConflict, sub.ID)
	}

	sub.Version = uint64(*newVersion)
	if updatedAt != nil {
		sub.UpdatedAt = *updatedAt
	}
	return nil
}

// Delete xóa subscription với CAS theo expectedVersion.
func (r *PostgresAmfSubRepo) Delete(ctx context.Context, id string, expectedVersion uint64) error {
	const query = `
		WITH target AS (
			SELECT subscription_id, version
			FROM amf_subscriptions
			WHERE subscription_id = $1
		), deleted AS (
			DELETE FROM amf_subscriptions s
			USING target t
			WHERE s.subscription_id = t.subscription_id AND t.version = $2
			RETURNING s.subscription_id
		)
		SELECT
			(SELECT count(*) FROM target)  AS found,
			(SELECT count(*) FROM deleted) AS deleted`

	var found, deleted int64
	if err := r.pool.QueryRow(ctx, query, id, int64(expectedVersion)).Scan(&found, &deleted); err != nil {
		return fmt.Errorf("delete subscription: %w", err)
	}

	switch {
	case found == 0:
		return fmt.Errorf("%w: subscription %s", domain.ErrNotFound, id)
	case deleted == 0:
		return fmt.Errorf("%w: subscription %s was modified concurrently", domain.ErrConflict, id)
	}
	return nil
}

// rowScanner khớp cả pgx.Row và pgx.Rows.
type rowScanner interface {
	Scan(dest ...any) error
}

func scanSubscription(row rowScanner) (*domain.AmfSubscription, error) {
	var (
		sub     domain.AmfSubscription
		version int64
		status  string
		trigger string
		data    []byte
	)

	err := row.Scan(
		&sub.ID, &version, &status, &data,
		&sub.EventNotifyUri, &sub.NotifyCorrelationId, &sub.NfId, &sub.Supi,
		&trigger, &sub.MaxReports, &sub.RemainReports, &sub.Expiry,
		&sub.CreatedAt, &sub.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	var payload model.AmfEventSubscription
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, fmt.Errorf("unmarshal subscription_data: %w", err)
	}

	sub.Version = uint64(version)
	sub.Status = domain.SubscriptionStatus(status)
	sub.TriggerType = model.AmfEventTrigger(trigger)
	sub.Data = payload
	return &sub, nil
}
