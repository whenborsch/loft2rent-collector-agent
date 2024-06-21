package loft

import (
	"github.com/whenborsch/loft2rent-collector-agent/internal/app/repository"
	"github.com/whenborsch/loft2rent-collector-agent/pkg/logger"
)

type Loft struct {
	repo repository.Loft

	log *logger.Logger
}

func New(repo repository.Loft) *Loft {
	return &Loft{
		repo: repo,
		log:  logger.GetInstance(),
	}
}
