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

// Get by id
func (r *ProductsRepository) FindById(id int) (product internal.Product, err error) {
	query := "SELECT * FROM products WHERE id = ?"
	row := r.DB.QueryRow(query, id)
	if row.Err() != nil {
		return internal.Product{}, row.Err()
	}
	err = row.Scan(&product.Id, &product.Name, &product.Quantity, &product.Code_value, &product.Is_published, &product.Expiration, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return internal.Product{}, internal.ErrNotFound
		}
		return internal.Product{}, err
	}
	return product, nil
}

// Create
func (r *ProductsRepository) Create(product internal.Product) (internal.Product, error) {
	query := "INSERT INTO products (name, quantity, code_value, is_published, expiration, price) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := r.DB.Exec(query, product.Name, product.Quantity, product.Code_value, product.Is_published, product.Expiration, product.Price)
	if err != nil {
		return internal.Product{}, err
	}
	lastInsertID, _ := result.LastInsertId()
	product.Id = int(lastInsertID)
	return product, nil
}

// Update
func (r *ProductsRepository) Update(product internal.Product) (internal.Product, error) {
	query := "UPDATE products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ? WHERE id = ?"
	_, err := r.DB.Exec(query, product.Name, product.Quantity, product.Code_value, product.Is_published, product.Expiration, product.Price, product.Id)
	if err != nil {
		return internal.Product{}, err
	}
	return product, nil
}

// Delete
func (r *ProductsRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
