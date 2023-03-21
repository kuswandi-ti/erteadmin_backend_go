package lingkungan

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(lingkungan Lingkungan) (Lingkungan, error)
	FindByNama(nama string) (Lingkungan, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(lingkungan Lingkungan) (Lingkungan, error) {
	err := r.db.Create(&lingkungan).Error
	if err != nil {
		return lingkungan, err
	}

	return lingkungan, nil
}

func (r *repository) FindByNama(nama string) (Lingkungan, error) {
	var lingkungan Lingkungan

	err := r.db.Where("nama = ?", nama).Find(&lingkungan).Error
	if err != nil {
		return lingkungan, err
	}

	return lingkungan, nil
}
