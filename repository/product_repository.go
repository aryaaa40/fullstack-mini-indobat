package repository

import (
	"mini-indobat-backend/database"
	"mini-indobat-backend/entity"
)

type ProductRepository interface {
	FindAll() ([]entity.Product, error)
	Create(product *entity.Product) error
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (r *productRepository) FindAll() ([]entity.Product, error) {
	var products []entity.Product
	err := database.DB.Find(&products).Error
	return products, err
}

func (r *productRepository) Create(product *entity.Product) error {
	return database.DB.Create(product).Error
}
