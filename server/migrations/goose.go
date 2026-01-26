package migrations

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

//go:embed sql/*.sql
var FS embed.FS

func Migrate(ctx context.Context, db *sql.DB) error {
	goose.SetBaseFS(FS)
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("goose: failed to set dialect: %w", err)
	}
	if err := goose.Up(db, "sql"); err != nil {
		return err
	}
	return nil
}
