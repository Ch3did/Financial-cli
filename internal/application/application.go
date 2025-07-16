package application

import (
	"financial-cli/internal/domain/category"
	"financial-cli/internal/domain/register"
	"financial-cli/internal/domain/transactions"
)

type Application struct {
	OFXImporter *OFXImporter
}

func NewApplication(
	categoryRepo *category.CategoryRepository,
	transactionRepo *transactions.TransactionRepository,
	registerRepo *register.RegisterRepository,
) *Application {
	return &Application{
		OFXImporter: NewOFXImporter(categoryRepo, transactionRepo, registerRepo),
	}
}
