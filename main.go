package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/rgalicia0729/database-go/pkg/product"
	"github.com/rgalicia0729/database-go/pkg/storage"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("Error al leer el archivo .env")
	}

	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	err := serviceProduct.Delete(3)
	if err != nil {
		fmt.Printf("product.Delete %v\n", err)
	}

}
