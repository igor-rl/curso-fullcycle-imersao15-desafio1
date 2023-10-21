package model

import (
	"encoding/binary"
	"errors"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type IProductRepository interface {
	RegisterProduct(product *Product) (*Product, error)
	ListProducts() ([]*Product, error)
	FindProductById(id string) (*Product, error)
}

type Product struct {
	ID          int32   `json:"id" gorm:"primaryKey;autoIncrement" valid:"-"`
	Name        string  `json:"name" gorm:"type:varchar(20);not null" valid:"notnull"`
	Description string  `json:"description" gorm:"type:varchar(255)" valid:"-"`
	Price       float64 `json:"price" gorm:"type:float;not null" valid:"notnull"`
}

func (product *Product) isValid() error {
	_, err := govalidator.ValidateStruct(product)

	if product.Price < 0.01 {
		return errors.New("o valor do produto precisa ser maior que 0.01")
	}

	if err != nil {
		return err
	}
	return nil
}

func NewProduct(name string, description string, price float64) (*Product, error) {
	product := Product{
		ID:          generateID(),
		Name:        name,
		Description: description,
		Price:       price,
	}

	err := product.isValid()
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func generateID() int32 {
	return int32(binary.BigEndian.Uint32(uuid.NewV4().Bytes()))
}
