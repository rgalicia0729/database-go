package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rgalicia0729/database-go/pkg/invoiceheader"
	"github.com/rgalicia0729/database-go/pkg/invoiceitem"
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

	if err := serviceProduct.Migrate(); err != nil {
		log.Fatalf("product.Migrate: %v\n", err)
	}

	storageInvoiceheader := storage.NewPsqlInvoiceheader(storage.Pool())
	serviceInvoiceheader := invoiceheader.NewService(storageInvoiceheader)

	if err := serviceInvoiceheader.Migrate(); err != nil {
		log.Fatalf("invoiceheader.Migrate: %v\n", err)
	}

	storageInvoiceitem := storage.NewPsqlInvoiceitem(storage.Pool())
	serviceInvoiceitem := invoiceitem.NewService(storageInvoiceitem)

	if err := serviceInvoiceitem.Migrate(); err != nil {
		log.Fatalf("invoiceitem.Migrate: %v\n", err)
	}
}
