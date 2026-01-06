package routes

import (
	controller "mini-indobat-backend/controller"
	"mini-indobat-backend/repository"
	"mini-indobat-backend/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	r := gin.Default()

	productRepo := repository.NewProductRepository()
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	orderRepo := repository.NewOrderRepository()
	orderService := service.NewOrderService(orderRepo, productRepo)
	orderController := controller.NewOrderController(orderService)

	r.GET("/products", productController.GetProducts)
	r.POST("/products", productController.CreateProduct)

	r.POST("/orders", orderController.CreateOrder)

	return r
}
