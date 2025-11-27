package postgres

import (
	"cashin-hexagonal/internal/domain/model"
	"cashin-hexagonal/internal/domain/repository"
	"context"

	"gorm.io/gorm"
)

type cashInRepository struct {
	db *gorm.DB
}

func NewCashInRepository(db *gorm.DB) repository.CashInRepository {
	return &cashInRepository{db: db}
}

func (r *cashInRepository) Save(ctx context.Context, cashin *model.CashIn) (string, error) {
	if err := r.db.WithContext(ctx).Create(cashin).Error; err != nil {
		return "", err
	}
	return cashin.ID.String(), nil
}

func (r *cashInRepository) FindByID(ctx context.Context, id string) (*model.CashIn, error) {
	var cashin model.CashIn
	if err := r.db.WithContext(ctx).First(&cashin, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &cashin, nil
}

// Para que GORM genere UUIDs
type CashIn = model.CashIn
