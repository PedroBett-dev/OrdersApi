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

type clientRepository interface {
	GetAll() ([]entities.Client, error)
	Create(client entities.Client) (entities.Client, error)
}

type ClientService struct {
	repo clientRepository
}

/*--------------------------------------------------------*/
func NewClientService(repo clientRepository) *ClientService {
	return &ClientService{repo: repo}
}

/*--------------------------------------------------------*/
func (s *ClientService) GetAll() ([]entities.Client, error) {
	return s.repo.GetAll()
}

/*--------------------------------------------------------*/
func (s *ClientService) Create(client entities.Client) (entities.Client, error) {

	if !etc.IsEmailValid(client.Email) {
		return entities.Client{}, apperror.ErrInvalidEmail
	}

	client.Cpf = etc.OnlyDigits(client.Cpf)
	if !brdoc.IsCPF(client.Cpf) {
		return entities.Client{}, apperror.ErrInvalidCpf
	}

	client.PhoneNumber = etc.OnlyDigits(client.PhoneNumber)
	if !etc.IsPhoneValid(client.PhoneNumber) {
		return entities.Client{}, apperror.ErrInvalidPhoneNumber
	}

	client.Address.State = strings.ToUpper(client.Address.State)
	uf, ok := etc.ParseUF(client.Address.State)
	if !ok {
		return entities.Client{}, apperror.ErrInvalidState
	}

	client.Address.Cep = etc.OnlyDigits(client.Address.Cep)
	if !brdoc.IsCEP(client.Address.Cep, uf) {
		return entities.Client{}, apperror.ErrInvalidCep
	}

	client.Id = etc.NewId(client.Cpf, idLength)

	created, err := s.repo.Create(client)
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return entities.Client{}, apperror.ErrClientAlreadyExists
	}

	return created, err
}

/*--------------------------------------------------------*/
