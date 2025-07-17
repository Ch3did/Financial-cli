package transactions

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	*gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

func (r *TransactionRepository) CreateIfNotExists(transaction *Transaction) error {
	var existing Transaction

	err := r.Where("value = ? AND date = ? AND description = ? AND transaction_id = ?",
		transaction.Value,
		transaction.Date,
		transaction.Description,
		transaction.TransactionID,
	).First(&existing).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return r.Create(transaction).Error
	}

	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepository) GetTopTransactionList(limit int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.Preload("Category").
		Order("date DESC").
		Limit(limit).
		Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *TransactionRepository) GetTransactionsInMonth(year int, month time.Month) ([]Transaction, error) {
	var transactions []Transaction

	start := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)

	if err := r.Where("date >= ? AND date < ? AND value < 0", start, end).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *TransactionRepository) SearchTransaction(value float64, date time.Time, description string, transactionID string) (*Transaction, error) {

	var existing Transaction

	err := r.Where("value = ? AND date = ? AND description = ? AND transaction_id = ?",
		value,
		date,
		description,
		transactionID,
	).
		First(&existing).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &existing, nil
}
