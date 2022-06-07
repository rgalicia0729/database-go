package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

// NewPostgresDB establish a connection to the database
func NewPostgresDB() {
	once.Do(func() {
		var err error

		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")

		connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("Can't open db: %v\n", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Can't do ping: %v\n", err)
		}

		log.Println("Conectado a la base de datos")
	})
}

// Pool return a unique instance of DB
func Pool() *sql.DB {
	return db
}
