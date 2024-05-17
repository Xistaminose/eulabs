package config

import (
	"database/sql"
	"log"
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
		db = InitPostgresDB()
	default:
		db = InitSQLiteDB()
	}

	if err != nil {
		return nil, "", err
	}

	return db, dbType, nil
}

func InitSQLiteDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./eulabs.db")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(sqliteSchema); err != nil {
		log.Fatal(err)
	}
	return db
}

const sqliteSchema = `
CREATE TABLE IF NOT EXISTS products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    price REAL,
    quantity INTEGER
);
`

func InitPostgresDB() *sql.DB {
	connStr := "user=postgres password=postgres dbname=eulabs sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(postgresSchema); err != nil {
		log.Fatal(err)
	}
	return db
}

const postgresSchema = `
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name TEXT,
    price REAL,
    quantity INTEGER
);
`
