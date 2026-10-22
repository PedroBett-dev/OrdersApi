package service

import (
	"errors"
	"strings"

	"github.com/paemuri/brdoc"
	"gorm.io/gorm"

	"github.com/seu-usuario/minha-api/internal/apperror"
	"github.com/seu-usuario/minha-api/internal/entities"
	"github.com/seu-usuario/minha-api/internal/etc"
)

const (
	idLength = 6
)

type customerRepository interface {
	GetAll() ([]entities.Customer, error)
	GetById(id string) (entities.Customer, error)
	Create(customer entities.Customer) (entities.Customer, error)
}

type CustomerService struct {
	repo customerRepository
}

func NewCustomerService(repo customerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) GetAll() ([]entities.Customer, error) {
	return s.repo.GetAll()
}

func (s *CustomerService) GetById(id string) (entities.Customer, error) {
	customer, err := s.repo.GetById(id)
	if err != nil {
		return entities.Customer{}, err
	}
	return customer, nil
}

func (s *CustomerService) GetAddress(id string) (entities.Address, error) {
	customer, err := s.repo.GetById(id)
	if err != nil {
		return entities.Address{}, err
	}
	return customer.Address, nil
}

func (s *CustomerService) Create(customer entities.Customer) (entities.Customer, error) {

	if !etc.IsEmailValid(customer.Email) {
		return entities.Customer{}, apperror.ErrInvalidEmail
	}

	customer.Cpf = etc.OnlyDigits(customer.Cpf)
	if !brdoc.IsCPF(customer.Cpf) {
		return entities.Customer{}, apperror.ErrInvalidCpf
	}

	customer.PhoneNumber = etc.OnlyDigits(customer.PhoneNumber)
	if !etc.IsPhoneValid(customer.PhoneNumber) {
		return entities.Customer{}, apperror.ErrInvalidPhoneNumber
	}

	customer.Address.State = strings.ToUpper(customer.Address.State)
	uf, ok := etc.ParseUF(customer.Address.State)
	if !ok {
		return entities.Customer{}, apperror.ErrInvalidState
	}

	customer.Address.Cep = etc.OnlyDigits(customer.Address.Cep)
	if !brdoc.IsCEP(customer.Address.Cep, uf) {
		return entities.Customer{}, apperror.ErrInvalidCep
	}

	customer.Status = "active"
	customer.Id = etc.NewId(customer.Cpf, idLength)

	created, err := s.repo.Create(customer)
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return entities.Customer{}, apperror.ErrCustomerAlreadyExists
	}

	return created, err
}
