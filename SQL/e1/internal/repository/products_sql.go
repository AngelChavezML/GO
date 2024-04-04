package repository

import (
	"database/sql"
	"go/sql/internal"
)

type ProductsRepository struct {
	DB *sql.DB
}

func (r *ProductsRepository) FindAll() (all_Products []internal.Product, err error) {
	rows, err := r.DB.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product internal.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Quantity, &product.Code_value, &product.Is_published, &product.Expiration, &product.Price)
		if err != nil {
			return nil, err
		}
		all_Products = append(all_Products, product)
	}

	return all_Products, nil
}
