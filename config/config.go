package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

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
