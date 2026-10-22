package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/seu-usuario/minha-api/internal/apperror"
	"github.com/seu-usuario/minha-api/internal/entities"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) GetAll() ([]entities.Customer, error) {
	customers := []entities.Customer{}
	if err := r.db.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *CustomerRepository) GetById(id string) (entities.Customer, error) {
	customer := entities.Customer{}
	if err := r.db.First(&customer, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Customer{}, apperror.ErrCustomerNotFound
		}
		return entities.Customer{}, err
	}
	return customer, nil
}

func (r *CustomerRepository) Create(customer entities.Customer) (entities.Customer, error) {
	if err := r.db.Create(&customer).Error; err != nil {
		return entities.Customer{}, err
	}
	return customer, nil
}
