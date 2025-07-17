package view

import (
	"financial-cli/internal/domain/category"
	"financial-cli/internal/domain/transactions"
	"fmt"
	"math"
	"strings"
)

func ShowSpendSummary(categories []category.Category, txs []transactions.Transaction) {
	RunIfNotDebug(ClearScreen)
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
}
