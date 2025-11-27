package model

import (
    "time"
    "github.com/google/uuid"
    "github.com/shopspring/decimal"
)

type CashIn struct {
    ID          uuid.UUID       `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
    AccountID   string          `gorm:"not null;index"`
    Amount      decimal.Decimal `gorm:"type:decimal(20,8);not null"`
    Currency    string          `gorm:"size:3;not null"`
    Reference   string          `gorm:"size:100"`
    Status      string          `gorm:"size:20;default:'PENDING'"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
