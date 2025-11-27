package repository

import (
	"cashin-hexagonal/internal/domain/model"
	"context"
)

type CashInRepository interface {
	Save(ctx context.Context, cashin *model.CashIn) (string, error)
	FindByID(ctx context.Context, id string) (*model.CashIn, error)
}
