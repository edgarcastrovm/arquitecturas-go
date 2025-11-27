package persistence

import (
	"cashin-modular/internal/cashin/domain/entity"
	"cashin-modular/internal/cashin/domain/repository"
	"context"

	"gorm.io/gorm"
)

type cashInRepository struct {
	db *gorm.DB
}

func NewCashInRepository(db *gorm.DB) repository.CashInRepository {
	db.AutoMigrate(&entity.CashIn{})
	return &cashInRepository{db: db}
}

func (r *cashInRepository) Save(ctx context.Context, cashin *entity.CashIn) error {
	return r.db.WithContext(ctx).Create(cashin).Error
}

func (r *cashInRepository) FindByID(ctx context.Context, id string) (*entity.CashIn, error) {
	var cashin entity.CashIn
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&cashin).Error
	if err != nil {
		return nil, err
	}
	return &cashin, nil
}

type CashIn = entity.CashIn
