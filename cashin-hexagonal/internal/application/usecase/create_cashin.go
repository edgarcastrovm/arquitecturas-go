package usecase

import (
	"cashin-hexagonal/internal/application/dto"
	"cashin-hexagonal/internal/domain/model"
	"cashin-hexagonal/internal/domain/repository"
	"context"
	"time"
)

type CreateCashInUseCase struct {
	repo repository.CashInRepository
}

func NewCreateCashInUseCase(repo repository.CashInRepository) *CreateCashInUseCase {
	return &CreateCashInUseCase{repo: repo}
}

func (uc *CreateCashInUseCase) Execute(ctx context.Context, req dto.CreateCashInRequest) (string, error) {
	cashin := &model.CashIn{
		AccountID: req.AccountID,
		Amount:    req.Amount,
		Currency:  req.Currency,
		Reference: req.Reference,
		Status:    "PENDING",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return uc.repo.Save(ctx, cashin)
}
