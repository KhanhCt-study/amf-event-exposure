// Package migrations nhúng các file SQL migration vào binary và áp dụng chúng
// trong một transaction duy nhất (READ COMMITTED mặc định của PostgreSQL).
package migrations

import (
	"context"
	"embed"
	"fmt"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed *.sql
var files embed.FS

// Up áp dụng lần lượt các file *.up.sql theo thứ tự tên file.
func Up(ctx context.Context, pool *pgxpool.Pool) error {
	if pool == nil {
		return fmt.Errorf("migrations: pool must not be nil")
	}
	names, err := scripts(".up.sql")
	if err != nil {
		return err
	}

	return pgx.BeginTxFunc(ctx, pool, pgx.TxOptions{IsoLevel: pgx.ReadCommitted}, func(tx pgx.Tx) error {
		for _, name := range names {
			content, err := files.ReadFile(name)
			if err != nil {
				return fmt.Errorf("migrations: read %s: %w", name, err)
			}
			if _, err := tx.Exec(ctx, string(content)); err != nil {
				return fmt.Errorf("migrations: apply %s: %w", name, err)
			}
		}
		return nil
	})
}

func scripts(suffix string) ([]string, error) {
	entries, err := files.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("migrations: read embedded dir: %w", err)
	}
	var names []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), suffix) {
			names = append(names, e.Name())
		}
	}
	sort.Strings(names)
	return names, nil
}
