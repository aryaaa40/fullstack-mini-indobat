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

	var req struct {
		ProductID       uint    `json:"product_id" binding:"required"`
		Quantity        int     `json:"quantity" binding:"required,gt=0"`
		DiscountPercent float64 `json:"discount_percent"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request input!",
			"details": err.Error(),
		})
		return
	}

	order, err := c.service.CreateOrder(req.ProductID, req.Quantity, req.DiscountPercent)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully!",
		"data":    order,
	})
}
