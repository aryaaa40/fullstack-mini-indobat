package controller

import (
	"net/http"

	"mini-indobat-backend/service"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service service.ProductService
}

func NewProductController(service service.ProductService) *ProductController {
	return &ProductController{
		service: service,
	}
}

func (c *ProductController) GetProducts(ctx *gin.Context) {

	products, err := c.service.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {

	var req struct {
		Name  string  `json:"name"`
		Stock int     `json:"stock"`
		Price float64 `json:"price"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body!"})
		return
	}

	product, err := c.service.CreateProduct(req.Name, req.Stock, req.Price)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}
