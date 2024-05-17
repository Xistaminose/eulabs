package config

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, string, error) {
	dbType := os.Getenv("DB_TYPE")
	var db *sql.DB
	var err error

	switch dbType {
	case "postgres":
		connStr := os.Getenv("POSTGRES_CONN_STR")
		if connStr == "" {
			connStr = "user=postgres password=postgres dbname=eulabs sslmode=disable"
		}
		db, err = sql.Open("postgres", connStr)
	default:
		db, err = sql.Open("sqlite3", "./eulabs.db")
	}

	if err != nil {
		return nil, "", err
	}

	return db, dbType, nil
}
