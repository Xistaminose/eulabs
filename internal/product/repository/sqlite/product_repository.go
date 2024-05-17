package sqlite

import (
	"database/sql"
	"eulabs/internal/entity"

	"log/slog"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) Create(product *entity.Product) (*entity.Product, error) {
	p, err := r.db.Exec("INSERT INTO products (name, price, quantity) VALUES (?, ?, ?)", product.Name, product.Price, product.Quantity)
	if err != nil {
		slog.Error("error creating product", err)
		return nil, err
	}

	id, err := p.LastInsertId()
	if err != nil {
		slog.Error("error getting last insert id", err)
		return nil, err
	}

	product.ID = int(id)
	return product, nil
}

func (r *ProductRepository) Fetch() ([]*entity.Product, error) {
	rows, err := r.db.Query("SELECT id, name, price, quantity FROM products ORDER BY id")
	if err != nil {
		slog.Error("error fetching products", err)
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		product := new(entity.Product)
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepository) GetByID(id int) (*entity.Product, error) {
	product := new(entity.Product)
	err := r.db.QueryRow("SELECT id, name, price, quantity FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
	if err != nil {
		slog.Error("error getting product by id", err)
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) Update(product *entity.Product) error {
	_, err := r.db.Exec("UPDATE products SET name = ?, price = ?, quantity = ? WHERE id = ?", product.Name, product.Price, product.Quantity, product.ID)
	if err != nil {
		slog.Error("error updating product", err)
	}
	return err
}

func (r *ProductRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		slog.Error("error deleting product", err)
	}
	return err
}
