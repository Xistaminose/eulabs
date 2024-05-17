package repository

import (
	"database/sql"
	"eulabs/internal/entity"
	"eulabs/internal/product/repository/postgresql"
	"eulabs/internal/product/repository/sqlite"
)

type ProductRepository interface {
	Create(product *entity.Product) (*entity.Product, error)
	Fetch() ([]*entity.Product, error)
	GetByID(id int) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id int) error
}

func NewProductRepository(db *sql.DB, dbType string) (ProductRepository, error) {
	switch dbType {
	case "postgres":
		return postgresql.NewProductRepository(db), nil
	default:
		return sqlite.NewProductRepository(db), nil
	}
}
