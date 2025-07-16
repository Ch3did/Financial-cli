package options

import (
	"financial-cli/internal/cli/view"
	"financial-cli/internal/domain/category"
	"fmt"

	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

func AddCategoryCommand(db *gorm.DB) *cli.Command {
	return &cli.Command{
		Name:  "add-category",
		Usage: "Adiciona uma nova categoria manualmente",
		Action: func(c *cli.Context) error {
			cat, err := view.PromptNewCategory()
			if err != nil {
				return fmt.Errorf("erro ao criar categoria: %w", err)
			}

			repo := category.NewCategoryRepository(db)

			if err := repo.Create(cat); err != nil {
				return fmt.Errorf("erro ao salvar categoria: %w", err)
			}

			entry := map[string]interface{}{
				"ID":       cat.ID,
				"Nome":     cat.Name,
				"Esperado": fmt.Sprintf("%.2f", cat.Expected),
			}

			view.BaseOutput(entry)

			return nil

		},
	}
}

func ListCategoriesCommand(db *gorm.DB) *cli.Command {
	return &cli.Command{
		Name:  "categories",
		Usage: "Lista todas as categorias registradas",
		Action: func(c *cli.Context) error {
			repo := category.NewCategoryRepository(db)

			categories, err := repo.GetAll()
			if err != nil {
				return fmt.Errorf("erro ao buscar categorias: %w", err)
			}

			var entries []map[string]interface{}

			for _, cat := range categories {
				entry := map[string]interface{}{
					"ID":        cat.ID,
					"Nome":      cat.Name,
					"Descrição": cat.Description,
					"Esperado":  fmt.Sprintf("%.2f", cat.Expected),
				}
				entries = append(entries, entry)
			}

			view.BaseOutput(entries)
			return nil
		},
	}
}
