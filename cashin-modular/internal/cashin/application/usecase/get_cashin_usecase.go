package usecase

import (
	"cashin-modular/internal/cashin/domain/entity"
	"cashin-modular/internal/cashin/domain/repository"
	"context"
)

type GetCashInUseCase struct {
	repo repository.CashInRepository
}

func NewGetCashInUseCase(repo repository.CashInRepository) *GetCashInUseCase {
	return &GetCashInUseCase{repo: repo}
}

func (uc *GetCashInUseCase) Execute(ctx context.Context, id string) (*entity.CashIn, error) {
	return uc.repo.FindByID(ctx, id)
}
