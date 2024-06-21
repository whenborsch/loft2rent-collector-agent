package sqlite

import (
	"database/sql"
	"github.com/whenborsch/loft2rent-collector-agent/config"
)

func New(cfg config.DatabaseConfig) (*sql.DB, error) {
	db, err := connect(cfg.Path)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connect(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}
