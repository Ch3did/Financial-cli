package options

import (
	"financial-cli/internal/cli/view"
	"financial-cli/internal/domain/category"
	"financial-cli/internal/domain/transactions"
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

func ShowCategoryTransactions(db *gorm.DB) *cli.Command {
	return &cli.Command{
		Name:  "transactions",
		Usage: "Busca todas as transações de uma categoria",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "category-id",
				Aliases: []string{"id"},
				Usage:   "ID da categoria",
				Value:   1,
			},
		},
		Action: func(c *cli.Context) error {
			view.RunIfNotDebug(view.ClearScreen)
			idStr := c.Path("id")
			idInt, err := strconv.Atoi(idStr)
			if err != nil {
				return fmt.Errorf("id inválido: %w", err)
			}

			id := uint(idInt)

			txRepo := transactions.NewTransactionRepository(db)
			catRepo := category.NewCategoryRepository(db)
			cat, err := catRepo.GetCategory(id)
			if err != nil {
				return fmt.Errorf("erro ao buscar categoria: %w", err)
			}
			txs, err := txRepo.GetCurrentMonthTransactionsByCategory(id)
			if err != nil {
				return fmt.Errorf("erro ao buscar transações do mês: %w", err)
			}
			var entry_list []interface{}
			for _, tx := range txs {
				entry := map[string]interface{}{
					"ID":              tx.ID,
					"Description":     tx.Description,
					"Date":            tx.Date,
					"Value":           tx.Value,
					"TransactionType": tx.TransactionType,
					"TransactionID":   tx.TransactionID,
					"Note":            tx.Note,
					"Category":        cat.Name,
				}
				entry_list = append(entry_list, entry)

			}

			view.BaseOutput(entry_list)
			return nil
		},
	}
}
