package application

import (
	"financial-cli/internal/domain/category"
	"financial-cli/internal/domain/register"
	"financial-cli/internal/domain/transactions"
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

func (i *OFXImporter) ImportAll(path string, txs []transactions.Transaction, reg register.Register) error {

	if err := i.RegisterRepo.Save(&reg).Error; err != nil {
		return err
	}

	for _, tx := range txs {
		tx.RegisterID = reg.ID
		if err := i.TransactionRepo.Save(&tx).Error; err != nil {
			return err
		}
	}

	return nil
}
