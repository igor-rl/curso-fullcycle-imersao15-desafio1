package usecase

import (
	"github.com/igorlage/curso/fullcycle/desafio1/domain/model"
)

type ProductUseCase struct {
	ProductRepository model.IProductRepository
}

func (p *ProductUseCase) RegisterProduct(name string, description string, price float64) (*model.Product, error) {
	product, err := model.NewProduct(name, description, price)
	if err != nil {
		return nil, err
	}
	p.ProductRepository.RegisterProduct(product)
	if product.ID < 0 {
		return nil, err
	}
	return product, nil
}

func (p *ProductUseCase) FindProducts() ([]*model.Product, error) {
	products, err := p.ProductRepository.ListProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductUseCase) FindProductById(id string) (*model.Product, error) {
	product, err := p.ProductRepository.FindProductById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
