package main

import (
	"ecommerce/api/controllers"
	services "ecommerce/products/application"
	products "ecommerce/products/domain/entities"
	shared "ecommerce/shared/Infrastructure"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()
	dsn := "host=localhost user=josias password=josias1228 dbname=ecommerce-GPT port=5432 sslmode=disable TimeZone=America/Santo_Domingo"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("error connecting to the database %s", err.Error()))
	}
	gormDB.AutoMigrate(&products.ProductDB{})

	productRepository := shared.NewGormRepository[*products.Product, *products.ProductDB](gormDB)
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
