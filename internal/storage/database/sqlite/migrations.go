package sqlite

import (
	"context"
	"database/sql"
	"embed"
	"github.com/pressly/goose/v3"
	"time"
)

func UpMigrations(fs embed.FS, db *sql.DB) error {
	const timeout = time.Minute * 2

	goose.SetBaseFS(fs)
	goose.SetLogger(goose.NopLogger())

	err := goose.SetDialect("sqlite")
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = goose.UpContext(ctx, db, "migrations")
	if err != nil {
		return err
	}

	return nil
}

func DownMigrations(fs embed.FS, db *sql.DB) error {
	const timeout = time.Minute * 2

	goose.SetBaseFS(fs)
	goose.SetLogger(goose.NopLogger())

	err := goose.SetDialect("sqlite")
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = goose.DownContext(ctx, db, "migrations")
	if err != nil {
		return err
	}

	return nil
}
