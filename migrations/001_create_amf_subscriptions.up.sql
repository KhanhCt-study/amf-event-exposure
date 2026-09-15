CREATE TABLE IF NOT EXISTS amf_subscriptions (
    subscription_id       UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    version               BIGINT NOT NULL CHECK (version > 0),
    status                TEXT NOT NULL DEFAULT 'active',

    -- Toàn bộ AmfEventSubscription dưới dạng JSONB
    subscription_data     JSONB NOT NULL,

    -- Cột denormalized phục vụ truy vấn
    event_notify_uri      TEXT NOT NULL,
    notify_correlation_id TEXT NOT NULL,
    nf_id                 TEXT NOT NULL,
    supi                  TEXT,

    -- Reporting mode
    trigger_type          TEXT NOT NULL DEFAULT 'CONTINUOUS',
    max_reports           INT,
    remain_reports        INT,
    expiry                TIMESTAMPTZ,

    created_at            TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
    updated_at            TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),

    CHECK (jsonb_typeof(subscription_data) = 'object'),
    CHECK (event_notify_uri IS NOT DISTINCT FROM subscription_data->>'eventNotifyUri'),
    CHECK (notify_correlation_id IS NOT DISTINCT FROM subscription_data->>'notifyCorrelationId'),
    CHECK (nf_id IS NOT DISTINCT FROM (subscription_data->>'nfId'))
);

CREATE INDEX IF NOT EXISTS amf_subscriptions_nf_id_idx
    ON amf_subscriptions (nf_id);
CREATE INDEX IF NOT EXISTS amf_subscriptions_supi_idx
    ON amf_subscriptions (supi) WHERE supi IS NOT NULL;
CREATE INDEX IF NOT EXISTS amf_subscriptions_notify_correlation_id_idx
    ON amf_subscriptions (notify_correlation_id);
