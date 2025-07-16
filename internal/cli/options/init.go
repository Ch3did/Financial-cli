package options

import (
	"financial-cli/internal/domain/category"
	"financial-cli/internal/domain/register"
	"financial-cli/internal/domain/transactions"
	"fmt"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

func InitDBCommand(db *gorm.DB) *cli.Command {
	return &cli.Command{
		Name:  "init",
		Usage: "Inicializa as tabelas no banco de dados",
		Action: func(c *cli.Context) error {
			err := db.AutoMigrate(
				&category.Category{},
				&transactions.Transaction{},
				&register.Register{},
			)
			if err != nil {
				return fmt.Errorf("erro ao inicializar banco: %w", err)
			}

			fmt.Println("Banco de dados inicializado com sucesso!")
			return nil
		},
	}
}
