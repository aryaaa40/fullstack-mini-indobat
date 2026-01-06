package repository

import (
	"mini-indobat-backend/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepository interface {
	GetProductForUpdate(tx *gorm.DB, productID uint) (*entity.Product, error)
	UpdateProductStock(tx *gorm.DB, product *entity.Product) error
	CreateOrder(tx *gorm.DB, order *entity.Order) error
}

type orderRepository struct{}

func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}

func (r *orderRepository) GetProductForUpdate(tx *gorm.DB, productID uint) (*entity.Product, error) {
	var product entity.Product
	err := tx.Clauses(
		clause.Locking{Strength: "UPDATE"},
	).First(&product, productID).Error

	return &product, err
}

func (r *orderRepository) UpdateProductStock(tx *gorm.DB, product *entity.Product) error {
	return tx.Save(product).Error
}

func (r *orderRepository) CreateOrder(tx *gorm.DB, order *entity.Order) error {
	return tx.Create(order).Error
}
