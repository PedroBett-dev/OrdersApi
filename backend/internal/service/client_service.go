package service

import (
	"crypto/sha256"
	"encoding/base32"
	"errors"

	"github.com/paemuri/brdoc"
	"gorm.io/gorm"

	"github.com/seu-usuario/minha-api/internal/apperror"
	"github.com/seu-usuario/minha-api/internal/entities"
)

const idLength = 6

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

	if !brdoc.IsCPF(client.Cpf) {
		return entities.Client{}, apperror.ErrInvalidCpf
	}

	client.Id = newId(client.Cpf)

	created, err := s.repo.Create(client)
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return entities.Client{}, apperror.ErrClientAlreadyExists
	}

	return created, err
}

/*--------------------------------------------------------*/
func newId(cpf string) string {
	sum := sha256.Sum256([]byte(cpf))
	return base32.StdEncoding.EncodeToString(sum[:])[:idLength]
}
