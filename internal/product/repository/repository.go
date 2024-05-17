package repository

import (
	"database/sql"
	"errors"
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
	case "sqlite3":
		return sqlite.NewProductRepository(db), nil
	default:
		return nil, errors.New("unsupported database type: " + dbType)
	}
}
