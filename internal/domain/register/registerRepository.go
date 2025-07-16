package register

import (
	"gorm.io/gorm"
)

type RegisterRepository struct {
	*gorm.DB
}

func NewRegisterRepository(db *gorm.DB) *RegisterRepository {
	return &RegisterRepository{db}
}

func (r *RegisterRepository) GetAll() ([]Register, error) {
	var Register []Register
	if err := r.Find(&Register).Error; err != nil {
		return nil, err
	}
	return Register, nil
}

func (r *RegisterRepository) Create(register *Register) (uint, error) {
	err := r.Save(register).Error
	return register.ID, err
}
