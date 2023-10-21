package repository

import (
	"github.com/igorlage/curso/fullcycle/desafio1/domain/model"
	"github.com/jinzhu/gorm"
)

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func (r ProductRepositoryDb) RegisterProduct(product *model.Product) (*model.Product, error) {
	err := r.Db.Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r ProductRepositoryDb) ListProducts() ([]*model.Product, error) {
	var products []*model.Product
	err := r.Db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r ProductRepositoryDb) FindProductById(id string) (*model.Product, error) {
	var product model.Product
	err := r.Db.First(&product, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
