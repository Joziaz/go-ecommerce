package main

import (
	"ecommerce/api/controllers"
	services "ecommerce/products/application"
	products "ecommerce/products/domain/entities"
	shared "ecommerce/shared/Infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	productRepository := shared.MemoryRepository[*products.Product]{}
	productService := services.NewProductService(&productRepository)
	productController := controllers.NewProductController(*productService)

	productsRoute := router.Group("products")
	{
		productsRoute.GET("/", productController.GetAll)
		productsRoute.GET("/:id", productController.GetById)
		productsRoute.POST("/", productController.SaveProduct)
		productsRoute.PUT("/", productController.Update)
		productsRoute.DELETE("/:id", productController.Delete)
	}

	router.Run()
}
