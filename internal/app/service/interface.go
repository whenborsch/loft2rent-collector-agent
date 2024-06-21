package service

import (
	"context"
	"github.com/whenborsch/loft2rent-collector-agent/internal/app/model"
)

type Loft interface {
	Create(ctx context.Context, loft model.Loft) error
	Get(ctx context.Context, loftID string) (*model.Loft, error)
	Update(ctx context.Context, loftID string, loft model.Loft) error
	Delete(ctx context.Context, loftID string) error
}
