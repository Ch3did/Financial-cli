package category

import (
	"fmt"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll() ([]Category, error) {
	var categories []Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) Create(cat *Category) error {
	if err := r.db.Create(cat).Error; err != nil {
		return fmt.Errorf("erro ao salvar categoria: %w", err)
	}
	return nil
}

func (r *CategoryRepository) GetCategory(id uint) (Category, error) {
	var cat Category
	if err := r.db.First(&cat, id).Error; err != nil {
		return cat, fmt.Errorf("erro ao buscar categoria com ID %d: %w", id, err)
	}
	return cat, nil
}
