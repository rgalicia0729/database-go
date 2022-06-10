package main

import (
	"database/sql"
	"errors"
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

	product, err := serviceProduct.GetById(5)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		log.Println("No hay un producto con este id")
	case err != nil:
		log.Fatalf("product.GetById: %v", err)
	default:
		fmt.Println(product)
	}

	// storageInvoiceheader := storage.NewPsqlInvoiceheader(storage.Pool())
	// serviceInvoiceheader := invoiceheader.NewService(storageInvoiceheader)

	// if err := serviceInvoiceheader.Migrate(); err != nil {
	// 	log.Fatalf("invoiceheader.Migrate: %v\n", err)
	// }

	// storageInvoiceitem := storage.NewPsqlInvoiceitem(storage.Pool())
	// serviceInvoiceitem := invoiceitem.NewService(storageInvoiceitem)

	// if err := serviceInvoiceitem.Migrate(); err != nil {
	// 	log.Fatalf("invoiceitem.Migrate: %v\n", err)
	// }
}
