package controller

import (
	"net/http"

	"mini-indobat-backend/service"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	service service.OrderService
}

func NewOrderController(service service.OrderService) *OrderController {
	return &OrderController{
		service: service,
	}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	// 1. Definisikan struct untuk request body
	var req struct {
		ProductID       uint    `json:"product_id" binding:"required"`
		Quantity        int     `json:"quantity" binding:"required,gt=0"`
		DiscountPercent float64 `json:"discount_percent"`
	}

	// 2. Bind JSON dari Postman ke struct req
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request input!",
			"details": err.Error(),
		})
		return
	}

	// 3. Panggil Service untuk memproses transaksi (Transaction & Locking terjadi di sini)
	order, err := c.service.CreateOrder(req.ProductID, req.Quantity, req.DiscountPercent)

	// 4. Handle jika ada error (Misal: Stok habis atau Product tidak ditemukan)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 5. Berikan respon sukses
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully!",
		"data":    order,
	})
}
