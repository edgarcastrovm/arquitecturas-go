package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CashInStatus string

const (
	Pending   CashInStatus = "pending"
	Completed CashInStatus = "completed"
	Failed    CashInStatus = "failed"
)

// type CashIn struct {
// 	ID        string       `gorm:"primaryKey"`
// 	AccountID string       `gorm:"index"`
// 	Amount    float64      `gorm:"type:decimal(10,2)"`
// 	Status    CashInStatus `gorm:"type:varchar(20)"`
// 	CreatedAt time.Time    `gorm:"autoCreateTime"`
// 	UpdatedAt time.Time    `gorm:"autoUpdateTime"`
// }

type CashIn struct {
	ID        uuid.UUID       `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	AccountID string          `gorm:"not null;index"`
	Amount    decimal.Decimal `gorm:"type:decimal(20,8);not null"`
	Currency  string          `gorm:"size:3;not null"`
	Reference string          `gorm:"size:100"`
	Status    string          `gorm:"size:20;default:'PENDING'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *CashIn) Validate() error {
	if c.Amount.Cmp(decimal.Zero) != 1 {
		return errors.New("amount must be positive")
	}
	if c.AccountID == "" {
		return errors.New("account_id is required")
	}
	return nil
}

func NewCashIn(id uuid.UUID, accountID string, amount decimal.Decimal) (*CashIn, error) {
	cashin := &CashIn{
		ID:        id,
		AccountID: accountID,
		Amount:    amount,
		Status:    string(Pending),
	}
	if err := cashin.Validate(); err != nil {
		return nil, err
	}
	return cashin, nil
}
