package Model

import "github.com/satori/go.uuid"

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewProduct() *Product {
	return &Product{
		ID: uuid.NewV4().String(),
	}
}

type Products struct {
	Product []*Product `json:"product"`
}

func (prod Products) Add(product *Product) {
	prod.Product = append(prod.Product, product)
}
