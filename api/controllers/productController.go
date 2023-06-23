package controllers

import (
	"net/http"

	. "ecommerce/products/application"
	. "ecommerce/products/domain/dtos"
	. "ecommerce/products/domain/entities"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductController struct {
	productService ProductService
}

func (controller *ProductController) GetAll(ctx *gin.Context) {
	products := controller.productService.GetAll()
	var dtos = make([]GetProductDto, len(products))

	for index, product := range products {
		dto := NewGetProductDto(*product)
		dtos[index] = dto
	}

	ctx.JSON(http.StatusOK, dtos)
}

func (controller *ProductController) GetById(ctx *gin.Context) {
	paramId := ctx.Params.ByName("id")
	id, err := uuid.Parse(paramId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
	}

	product, err := controller.productService.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	dto := NewGetProductDto(*product)

	ctx.JSON(http.StatusOK, dto)
}

func (controller *ProductController) SaveProduct(ctx *gin.Context) {
	var product = Product{}
	ctx.BindJSON(&product)

	if product.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "product name is required"})
	}

	if product.Description == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "product description is required"})
	}

	if product.Price <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "product price must be bigger than 0"})
	}

	savedProduct, err := controller.productService.Save(product)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	dto := NewGetProductDto(*savedProduct)
	ctx.JSON(http.StatusCreated, dto)
}

func (controller *ProductController) Update(ctx *gin.Context) {
	var product = Product{}
	ctx.BindJSON(&product)
	err := controller.productService.Update(product)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
}

func (controller *ProductController) Delete(ctx *gin.Context) {
	paramId := ctx.Params.ByName("id")
	id, err := uuid.Parse(paramId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = controller.productService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
}

func NewProductController(service ProductService) *ProductController {
	return &ProductController{service}
}
