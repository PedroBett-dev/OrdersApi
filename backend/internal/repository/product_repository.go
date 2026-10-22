package repository

import (
	"github.com/seu-usuario/minha-api/internal/entities"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]entities.Product, error) {
	products := []entities.Product{}
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
