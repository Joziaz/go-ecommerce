package products

import (
	products "ecommerce/products/domain/entities"
	shared "ecommerce/shared/domain"

	"github.com/google/uuid"
)

type UpdateProductDto struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (dto UpdateProductDto) ToProduct() (products.Product, error) {
	id, err := uuid.Parse(dto.Id)
	if err != nil {
		return products.Product{}, shared.NewDomainError("Invalid Id")
	}

	product := products.Product{
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
	}
	product.SetId(id)

	return product, nil
}

func (dto UpdateProductDto) Validate() []error {
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
