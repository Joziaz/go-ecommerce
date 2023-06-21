package products

import (
	products "ecommerce/products/domain/entities"
	shared "ecommerce/shared/domain/interfaces"

	"github.com/google/uuid"
)

type ProductService struct {
	productRepository shared.Repository[*products.Product]
}

func NewProductService(repo shared.Repository[*products.Product]) *ProductService {
	return &ProductService{repo}
}

func (service *ProductService) Save(product products.Product) *products.Product {
	return service.productRepository.Save(&product)
}

func (service *ProductService) GetById(id uuid.UUID) (*products.Product, error) {
	product, err := service.productRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service *ProductService) GetAll() []*products.Product {
	return service.productRepository.GetAll()
}

func (service *ProductService) Update(product products.Product) error {
	err := service.productRepository.Update(&product)
	if err != nil {
		return err
	}

	return nil
}

func (service *ProductService) Delete(id uuid.UUID) error {
	err := service.productRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
