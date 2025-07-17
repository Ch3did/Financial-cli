package application

import (
	"financial-cli/internal/domain/category"
	"financial-cli/internal/domain/register"
	"financial-cli/internal/domain/transactions"
	"time"
)

type OFXImporter struct {
	CategoryRepo    *category.CategoryRepository
	TransactionRepo *transactions.TransactionRepository
	RegisterRepo    *register.RegisterRepository
}

func NewOFXImporter(
	catRepo *category.CategoryRepository,
	txRepo *transactions.TransactionRepository,
	regRepo *register.RegisterRepository,
) *OFXImporter {
	return &OFXImporter{
		CategoryRepo:    catRepo,
		TransactionRepo: txRepo,
		RegisterRepo:    regRepo,
	}
}

func (i *OFXImporter) ImportTransaction(tx transactions.Transaction) error {
	err := i.TransactionRepo.CreateIfNotExists(&tx)
	if err != nil {
		return err
	}
	return nil
}

func (i *OFXImporter) CreateRegister(reg register.Register) (uint, error) {
	regId, err := i.RegisterRepo.Create(&reg)
	if err != nil {
		return 0, err
	}
	return regId, nil
}

func (i *OFXImporter) SearchDuplicateTransaction(value float64, date time.Time, description string, transactionID string) (*transactions.Transaction, error) {
	return i.TransactionRepo.SearchTransaction(value, date, description, transactionID)
}
