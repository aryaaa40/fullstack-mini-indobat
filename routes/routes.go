package routes

import (
	controller "mini-indobat-backend/handler"
	"mini-indobat-backend/repository"
	"mini-indobat-backend/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	r := gin.Default()

	productRepo := repository.NewProductRepository()
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	// routes
	r.GET("/products", productController.GetProducts)
	r.POST("/products", productController.CreateProduct)

	return r
}
