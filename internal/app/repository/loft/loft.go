package loft

import (
	"database/sql"
	"github.com/whenborsch/loft2rent-collector-agent/pkg/logger"
)

type Loft struct {
	db *sql.DB

	log *logger.Logger
}

func New(db *sql.DB) *Loft {
	return &Loft{
		db: db,
	}
}
