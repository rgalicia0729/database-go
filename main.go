package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rgalicia0729/database-go/pkg/storage"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("Error al leer el archivo .env")
	}

	storage.NewPostgresDB()
}
