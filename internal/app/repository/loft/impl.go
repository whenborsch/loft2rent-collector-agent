package loft

import (
	"context"
	"github.com/whenborsch/loft2rent-collector-agent/internal/app/model"
)

func (r *Loft) Create(ctx context.Context, loft ...model.Loft) error {
	return nil
}

func (r *Loft) Get(ctx context.Context, loftID string) (*model.Loft, error) {
	return nil, nil
}

func (r *Loft) Update(ctx context.Context, loftID string, loft model.Loft) error {
	return nil
}

func (r *Loft) Delete(ctx context.Context, loftID string) error {
	return nil
}
