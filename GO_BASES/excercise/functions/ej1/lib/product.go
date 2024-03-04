package lib

// Struct of product
type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products []Product

func (p Product) Save(Name string, Description string, Category string, ID int, Price float64) []Product {
	Products = append(Products, Product{ID: ID, Name: Name, Price: Price, Description: Description, Category: Category})
	return Products
}
func (p Product) GetAll() []Product {
	return Products
}
