package register

import (
	"time"

	"gorm.io/gorm"
)

type Register struct {
	ID           uint           `gorm:"primaryKey;autoIncrement"`
	OrgID        string         `gorm:"column:org_id;not null"`
	Account      string         `gorm:"not null"`
	StartDate    time.Time      `gorm:"column:start_date;not null"`
	EndDate      time.Time      `gorm:"column:end_date;not null"`
	Organization string         `gorm:"not null"`
	Amount       float64        `gorm:"not null"`
	CreatedAt    time.Time      `gorm:"column:creation_date;autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"column:change_timestamp;autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type RegisterInteface interface {
	GetAll() ([]Register, error)
	Create(register *Register) (uint, error)
}
