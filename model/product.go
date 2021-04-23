package model

import uuid "github.com/satori/go.uuid"

type Product struct {
	// o go tem algo chamado de tags e eu coloquei as tags para quando eu
	// trabalhar com json eu vou colocar minusculo
	ID   string `json: "id"`
	Name string `json: "name"`
}

// criando uma lista de produtos
type Products struct {
	Product []Product
}

// função que vai adicionar um produto na lista de produtos
func (p *Products) Add(product Product) {
	p.Product = append(p.Product, product)
}

// função que vai criar um novo produto
func NewProduct() *Product {
	return &Product{
		ID: uuid.NewV4().String(),
	}
}
