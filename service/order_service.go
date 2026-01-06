package service

import (
	"errors"
	"mini-indobat-backend/database"
	"mini-indobat-backend/entity"
	"mini-indobat-backend/repository"

	"gorm.io/gorm"
)

type OrderService interface {
	CreateOrder(productID uint, qty int, discount float64) (*entity.Order, error)
}

type orderService struct {
	orderRepo   repository.OrderRepository
	productRepo repository.ProductRepository
}

func NewOrderService(orderRepo repository.OrderRepository, productRepo repository.ProductRepository) OrderService {
	return &orderService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (s *orderService) CreateOrder(productID uint, qty int, discount float64) (*entity.Order, error) {
	if qty <= 0 {
		return nil, errors.New("Quantity must be greater than 0!")
	}

	if discount < 0 || discount > 100 {
		return nil, errors.New("Discount percent must be between 0–100!")
	}

	var order *entity.Order

	err := database.DB.Transaction(func(tx *gorm.DB) error {

		product, err := s.orderRepo.GetProductForUpdate(tx, productID)
		if err != nil {
			return errors.New("Product not found!")
		}

		if product.Stock < qty {
			return errors.New("Insufficient stock!")
		}

		total := product.Price * int64(qty)
		discountAmount := int64(float64(total) * (discount / 100))
		totalAfterDiscount := total - discountAmount

		product.Stock -= qty
		if err := s.orderRepo.UpdateProductStock(tx, product); err != nil {
			return err
		}

		newOrder := &entity.Order{
			ProductID:       productID,
			Quantity:        qty,
			DiscountPercent: discount,
			TotalAmount:     totalAfterDiscount,
		}

		if err := s.orderRepo.CreateOrder(tx, newOrder); err != nil {
			return err
		}

		order = newOrder
		return nil
	})

	return order, err
}
