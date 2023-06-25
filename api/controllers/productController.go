package controllers

import (
	"errors"
	"net/http"

	. "ecommerce/products/application"
	. "ecommerce/products/domain/dtos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductController struct {
	productService ProductService
}

func NewProductController(service ProductService) *ProductController {
	return &ProductController{service}
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
		ctx.Error(errors.New("invalid id, must be a valid UUID"))
		return
	}

	product, err := controller.productService.GetById(id)
	if err != nil {
		ctx.Error(err)
		return
	}

	dto := NewGetProductDto(*product)

	ctx.JSON(http.StatusOK, dto)
}

func (controller *ProductController) SaveProduct(ctx *gin.Context) {
	var productDto = SaveProductDto{}
	ctx.BindJSON(&productDto)

	errs := productDto.Validate()
	if errs != nil {
		for _, err := range errs {
			ctx.Error(err)
		}
	}

	if ctx.Errors != nil {
		return
	}

	savedProduct, err := controller.productService.Save(productDto.ToProduct())
	if err != nil {
		ctx.Error(err)
		return
	}

	dto := NewGetProductDto(*savedProduct)
	ctx.JSON(http.StatusCreated, dto)
}

func (controller *ProductController) Update(ctx *gin.Context) {
	var productDto = SaveProductDto{}
	ctx.BindJSON(&productDto)

	errs := productDto.Validate()
	if errs != nil {
		for _, err := range errs {
			ctx.Error(err)
		}
	}

	err := controller.productService.Update(productDto.ToProduct())
	if err != nil {
		ctx.Error(err)
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
		ctx.Error(err)
		return
	}
}
