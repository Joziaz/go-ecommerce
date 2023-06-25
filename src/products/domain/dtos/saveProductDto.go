package products

import (
	products "ecommerce/products/domain/entities"
	shared "ecommerce/shared/domain"
)

type SaveProductDto struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (dto SaveProductDto) ToProduct() products.Product {
	return products.Product{
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
	}
}

func (dto SaveProductDto) Validate() []error {
	var errorsList []error = nil

	if dto.Name == "" {
		errorsList = append(errorsList, shared.NewDomainError("product name can't be empty"))
	}

	if dto.Description == "" {
		errorsList = append(errorsList, shared.NewDomainError("product description can't be empty"))
	}

	if dto.Price <= 0 {
		errorsList = append(errorsList, shared.NewDomainError("product price must be bigger than 0"))
	}

	return errorsList
}
