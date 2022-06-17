package domain

import (
	"context"
	"errors"
)

var (
	ErrProductNotFound      = errors.New("Product not found")
	ErrProductStockNegative = errors.New("Negative inventory")
)

type (
	Size struct {
		large float64
		width float64
		deep  float64
	}

	Product struct {
		name       string
		cathegory  string
		price      float64
		dimension  Size
		visibility bool
		stock      int
	}

	ProductRepository interface {
		Create(context.Context, Product) (Product, error)
		FindById(context.Context, string) (Product, error)
		FindAll(context.Context) ([]Product, error)
		UpdateStock(context.Context, Product) error
		UpdateVisibility(context.Context, Product) error
	}
)

func NewSize(large float64, width float64, deep float64) Size {
	return Size{
		large: large,
		width: width,
		deep:  deep,
	}
}

func (s Size) Large() float64 {
	return s.large
}

func (s Size) Width() float64 {
	return s.width
}

func (s Size) Deep() float64 {
	return s.deep
}

func NewProduct(name string, cathegory string, price float64,
	visibility bool, large float64, width float64, deep float64) Product {

	s := NewSize(large, width, deep)

	return Product{
		name:       name,
		cathegory:  cathegory,
		price:      price,
		visibility: visibility,
		dimension:  s,
		stock:      0,
	}

}

func (p *Product) UpdateStock(increment int) error {
	p.stock = p.stock + increment
	if p.stock < 0 {
		return ErrProductStockNegative
	}
	return nil
}

func (p Product) Name() string {
	return p.name
}

func (p Product) Cathegory() string {
	return p.cathegory
}

func (p Product) Price() float64 {
	return p.price
}

func (p Product) Visibility() bool {
	return p.visibility
}

func (p Product) Dimencion() Size {
	return p.dimension
}

func (p Product) Stock() int {
	return p.stock
}
