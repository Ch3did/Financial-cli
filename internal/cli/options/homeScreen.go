package options

import (
	"fmt"
	"math"
	"strings"
	"time"

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

			spentByCategory := make(map[uint]float64)
			for _, tx := range txs {
				if tx.CategoryID != 0 {
					spentByCategory[tx.CategoryID] += tx.Value
				}
			}

			barLen := 45
			fmt.Println("Resumo de gastos do mês atual:")
			fmt.Println("-------------------------------------------------------------------------------")

			for _, cat := range categories {
				total := -cat.Expected
				spent := -spentByCategory[cat.ID]

				var filledLen int
				if total > 0 {
					percent := spent / total
					if percent > 1 {
						percent = 1
					}
					filledLen = int(float64(barLen) * percent)
				} else {
					filledLen = barLen
				}

				bar := strings.Repeat("#", filledLen) + strings.Repeat("-", barLen-filledLen)

				valorParaMostrar := math.Abs(spentByCategory[cat.ID])

				fmt.Printf(
					"%-15s: [%s]| Gasto: %.2f\n",
					cat.Name,
					bar,
					valorParaMostrar,
				)
			}
			return nil
		},
	}
}
