package options

import (
	"fmt"
	"time"

	"financial-cli/internal/cli/view"
	"financial-cli/internal/domain/category"
	"financial-cli/internal/domain/transactions"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

func ShowSpendSummaryCommand(db *gorm.DB) *cli.Command {
	return &cli.Command{
		Name:  "home",
		Usage: "Mostra resumo do gasto mensal com barras de progresso",
		Action: func(c *cli.Context) error {
			catRepo := category.NewCategoryRepository(db)
			txRepo := transactions.NewTransactionRepository(db)

			categories, err := catRepo.GetAll()
			if err != nil {
				return fmt.Errorf("erro ao buscar categorias de gasto: %w", err)
			}

			now := time.Now()
			year, month := now.Year(), now.Month()
			txs, err := txRepo.GetTransactionsInMonth(year, month)
			if err != nil {
				return fmt.Errorf("erro ao buscar transações do mês: %w", err)
			}

			view.ShowSpendSummary(categories, txs)

			return nil
		},
	}
}
