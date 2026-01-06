package service

import (
	"errors"
	"mini-indobat-backend/entity"
	"mini-indobat-backend/repository"
)

type ProductService interface {
	GetProducts() ([]entity.Product, error)
	CreateProduct(name string, stock int, price float64) (*entity.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) GetProducts() ([]entity.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) CreateProduct(name string, stock int, price float64) (*entity.Product, error) {

	if name == "" {
		return nil, errors.New("Product name is required!")
	}

	if stock < 0 {
		return nil, errors.New("Stock cannot be negative!")
	}

	if price <= 0 {
		return nil, errors.New("Price must be greater than zero!")
	}

	product := &entity.Product{
		Name:  name,
		Stock: stock,
		Price: int64(price),
	}

	err := s.repo.Create(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
