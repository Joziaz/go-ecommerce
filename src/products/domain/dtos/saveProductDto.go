package products

import products "ecommerce/products/domain/entities"

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
