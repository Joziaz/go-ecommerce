package products

import (
	interfaces "ecommerce/shared/domain/interfaces"
	shared "ecommerce/shared/domain/models"

	"github.com/google/uuid"
)

type ProductDB struct {
	shared.GormModel
	Name        string
	Description string
	Price       float64
}

func (p ProductDB) TableName() string {
	return "products"
}

func (p *ProductDB) ToEntity() *Product {
	product := Product{
		BaseEntity: shared.BaseEntity{
			CreatedAt: p.CreatedAt,
		},
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}
	product.SetId(uuid.MustParse(p.ID))

	return &product
}

func (p *ProductDB) FromEntity(product *Product) interfaces.EntityDB[*Product] {
	return &ProductDB{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		GormModel: shared.GormModel{
			ID: product.GetId().String(),
		},
	}
}
