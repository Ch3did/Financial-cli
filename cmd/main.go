package main

import (
	"financial-cli/internal/cli/options"
	"financial-cli/internal/config/database"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB()

	app := &cli.App{
		Name:  "financial-cli",
		Usage: "A command-line financial manager",
		Commands: []*cli.Command{
			options.ImportCommand(db),
			options.InitDBCommand(db),
			options.AddCategoryCommand(db),
			options.ListCategoriesCommand(db),
			options.ShowSpendSummaryCommand(db),
			options.ShowCategoryTransactions(db),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
