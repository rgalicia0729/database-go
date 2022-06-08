package storage

import (
	"database/sql"
	"log"
)

const (
	psqlMigrateProduct = `
		CREATE TABLE IF NOT EXISTS products(
			id SERIAL NOT NULL,
			name CHARACTER VARYING(25) NOT NULL,
			observations CHARACTER VARYING(100),
			price INT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP,
			CONSTRAINT products_id_pk PRIMARY KEY (id)
		)
	`
)

// PsqlProduct used for work with postgres - product
type PsqlProduct struct {
	db *sql.DB
}

// NewPsqlProduct return a new pointer of PsqlProduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	log.Println("Migración de producto ejecutada correctamente")

	return nil
}
