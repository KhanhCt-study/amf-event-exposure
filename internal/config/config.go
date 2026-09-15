// Package config đọc cấu hình service từ biến môi trường.
package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/KhanhCt-study/amf-event-exposure.git/internal/delivery/http/validator"
)

// Config gom toàn bộ tham số runtime của AMF Event Exposure Service.
type Config struct {
	ListenAddr      string
	PublicAPIURL    string
	NfInstanceID    string
	DatabaseURL     string
	HTTPTimeout     time.Duration
	DatabaseTimeout time.Duration
	ShutdownTimeout time.Duration

	// Vị trí mặc định dùng cho prototype LOCATION_REPORT.
	DefaultMcc      string
	DefaultMnc      string
	DefaultTac      string
	DefaultNrCellID string
}

// Load đọc config từ env và validate các field bắt buộc.
func Load() (*Config, error) {
	cfg := &Config{
		ListenAddr:      getEnv("AMF_LISTEN_ADDR", ":8081"),
		PublicAPIURL:    getEnv("AMF_PUBLIC_API_URL", "http://localhost:8081"),
		NfInstanceID:    os.Getenv("AMF_NF_INSTANCE_ID"),
		DatabaseURL:     os.Getenv("DATABASE_URL"),
		DefaultMcc:      getEnv("AMF_DEFAULT_MCC", "208"),
		DefaultMnc:      getEnv("AMF_DEFAULT_MNC", "95"),
		DefaultTac:      getEnv("AMF_DEFAULT_TAC", "000001"),
		DefaultNrCellID: getEnv("AMF_DEFAULT_NR_CELL_ID", "000000001"),
	}

	var err error
	if cfg.HTTPTimeout, err = getDuration("AMF_TIMEOUT", 10*time.Second); err != nil {
		return nil, err
	}
	if cfg.DatabaseTimeout, err = getDuration("DATABASE_TIMEOUT", 5*time.Second); err != nil {
		return nil, err
	}
	if cfg.ShutdownTimeout, err = getDuration("AMF_SHUTDOWN_TIMEOUT", 15*time.Second); err != nil {
		return nil, err
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *Config) validate() error {
	if strings.TrimSpace(c.ListenAddr) == "" {
		return fmt.Errorf("config: AMF_LISTEN_ADDR must not be empty")
	}
	if !validator.IsAbsoluteHTTPURL(c.PublicAPIURL) {
		return fmt.Errorf("config: AMF_PUBLIC_API_URL must be an absolute http(s) URL")
	}
	if strings.TrimSpace(c.DatabaseURL) == "" {
		return fmt.Errorf("config: DATABASE_URL is required")
	}
	if c.NfInstanceID != "" && !validator.IsUUID(c.NfInstanceID) {
		return fmt.Errorf("config: AMF_NF_INSTANCE_ID must be a valid UUID")
	}
	return nil
}

func getEnv(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

func getDuration(key string, fallback time.Duration) (time.Duration, error) {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return fallback, nil
	}
	d, err := time.ParseDuration(raw)
	if err != nil {
		return 0, fmt.Errorf("config: %s is not a valid duration: %w", key, err)
	}
	if d <= 0 {
		return 0, fmt.Errorf("config: %s must be greater than 0", key)
	}
	return d, nil
}
