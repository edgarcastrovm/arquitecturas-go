package usecase

import (
	"cashin-modular/internal/cashin/domain/entity"
	"cashin-modular/internal/cashin/domain/repository"
	"context"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateCashInUseCase struct {
	repo repository.CashInRepository
}

func NewCreateCashInUseCase(repo repository.CashInRepository) *CreateCashInUseCase {
	return &CreateCashInUseCase{repo: repo}
}

type CreateCashInOutput struct {
	ID      uuid.UUID `json:"id"`
	Message string    `json:"message"`
}

func (uc *CreateCashInUseCase) Execute(ctx context.Context, accountID string, amount decimal.Decimal) (*CreateCashInOutput, error) {
	id := uuid.New()
	cashin, err := entity.NewCashIn(id, accountID, amount)
	if err != nil {
		return nil, err
	}

	if err := uc.repo.Save(ctx, cashin); err != nil {
		return nil, err
	}

	// Aquí iría lógica de negocio: notificaciones push,email, sms, etc.
	// Por ejemplo: producer.SendCashInEvent(cashin)

	return &CreateCashInOutput{ID: id, Message: "CashIn registrado correctamente"}, nil
}
