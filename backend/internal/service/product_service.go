package service

import "github.com/seu-usuario/minha-api/internal/entities"

type productRepository interface {
	GetAll() ([]entities.Product, error)
}

type ProductService struct {
	repo productRepository
}

func NewProductService(repo productRepository) *ProductService { return &ProductService{repo: repo} }

func (s *ProductService) GetAll() ([]entities.Product, error) {
	return s.repo.GetAll()
}
