package products

import shared "ecommerce/shared/domain/models"

type Product struct {
	shared.BaseEntity
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
