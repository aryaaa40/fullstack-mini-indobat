package entity

import "time"

type Order struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	ProductID       uint      `gorm:"not null" json:"product_id"`
	Quantity        int       `gorm:"not null" json:"quantity"`
	DiscountPercent float64   `gorm:"default:0" json:"discount_percent"`
	TotalAmount     int64     `gorm:"not null" json:"total_amount"`
	CreatedAt       time.Time `json:"created_at"`
}
