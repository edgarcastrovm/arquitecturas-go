package dto

import "github.com/shopspring/decimal"

type CreateCashInRequest struct {
	AccountID string          `json:"account_id" binding:"required"`
	Amount    decimal.Decimal `json:"amount" binding:"required,gt=0"`
	Currency  string          `json:"currency" binding:"required,len=3"`
	Reference string          `json:"reference,omitempty"`
}
