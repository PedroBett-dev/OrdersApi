package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/seu-usuario/minha-api/internal/apperror"
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

/*--------------------------------------------------------*/
func (r *ClientRepository) GetById(id string) (entities.Client, error) {
	client := entities.Client{}
	if err := r.db.First(&client, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Client{}, apperror.ErrClientNotFound
		}
		return entities.Client{}, err
	}
	return client, nil
}

/*--------------------------------------------------------*/
func (r *ClientRepository) Create(client entities.Client) (entities.Client, error) {
	if err := r.db.Create(&client).Error; err != nil {
		return entities.Client{}, err
	}
	return client, nil
}
