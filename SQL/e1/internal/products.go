package internal

import "time"

type Product struct {
	Id           int
	Name         string
	Quantity     int
	Code_value   string
	Is_published bool
	Expiration   time.Time
	Price        float64
}
type ProductRepository interface {
	FindAll() ([]Product, error)
	FindById(id int) (Product, error)
	Create(product Product) (Product, error)
	Update(product Product) (Product, error)
	Delete(id int) error
}
