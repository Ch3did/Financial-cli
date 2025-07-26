package category

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Expected    float64   `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt
}

type CategoryInterface interface {
	GetAll() ([]Category, error)
	Create(cat *Category) error
	GetSpendCategories() ([]Category, error)
	GetCategory(id int) (*Category, error)
}
