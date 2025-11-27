package repository

import (
	"cashin-modular/internal/cashin/domain/entity"
	"context"
)

type CashInRepository interface {
	Save(ctx context.Context, cashin *entity.CashIn) error
	FindByID(ctx context.Context, id string) (*entity.CashIn, error)
	// List(ctx context.Context, filter Filter) ([]*entity.CashIn, error) // Futuro
}
