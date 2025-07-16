package transactions

import (
	"financial-cli/internal/domain/category"
	"financial-cli/internal/domain/register"
	"time"
)

type Transaction struct {
	ID              uint      `gorm:"primaryKey"`
	Description     string    `gorm:"not null"`
	Date            time.Time `gorm:"not null"`
	Value           float64   `gorm:"not null"`
	TransactionType string    `gorm:"not null"`
	TransactionID   string    `gorm:"not null"`
	Note            string    `gorm:"not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	CategoryID uint               `gorm:"index"`
	Category   *category.Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	RegisterID uint              `gorm:"index"`
	Register   register.Register `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type TransactionInterface interface {
	CreateOrUpdate() error
	GetTopTransactionList(limit int) ([]Transaction, error)
	GetTransactionsInMonth(year int, month time.Month) ([]Transaction, error)
}
