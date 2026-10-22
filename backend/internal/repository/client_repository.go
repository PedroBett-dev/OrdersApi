package repository

import (
	"gorm.io/gorm"

	"github.com/seu-usuario/minha-api/internal/entities"
)

type ClientRepository struct {
	db *gorm.DB
}

/*--------------------------------------------------------*/
func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

/*--------------------------------------------------------*/
func (r *ClientRepository) GetAll() ([]entities.Client, error) {
	clients := []entities.Client{}

	if err := r.db.Find(&clients).Error; err != nil {
		return nil, err
	}

	return clients, nil
}
