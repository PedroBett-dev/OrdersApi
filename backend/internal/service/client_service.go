package service

import "github.com/seu-usuario/minha-api/internal/entities"

type clientRepository interface {
	GetAll() ([]entities.Client, error)
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
