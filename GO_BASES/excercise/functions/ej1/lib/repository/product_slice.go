package repository

import "exc/func/lib"

var Products []lib.Product

func main() {
	Products = append(Products, lib.Product{ID: 1, Name: "Product 1", Price: 100, Description: "Description 1", Category: "Category 1"})
	//Use the function Save from lib/product.go
	Products = lib.Product.Save("Product 2", "Description 2", "Category 2", 200, 1000)
	//Use the function GetAll from lib/product.go
	Products = lib.Product.GetAll()

}
