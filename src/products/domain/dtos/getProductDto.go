package products

import (
	products "ecommerce/products/domain/entities"

	"github.com/google/uuid"
)

type GetProductDto struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
}

func NewGetProductDto(product products.Product) GetProductDto {
	return GetProductDto{
		Id:          product.GetId(),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}
}
