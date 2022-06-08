package storage

import (
	"database/sql"
	"log"
)

const (
	psqlMigrateInvoiceheader = `
		CREATE TABLE IF NOT EXISTS invoice_headers(
			id SERIAL NOT NULL,
			client CHARACTER VARYING(100) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP,
			CONSTRAINT invoice_headers_id_pk PRIMARY KEY (id)
		)
	`
)

// PsqlInvoiceheader used for work with postgres - invoiceheader
type PsqlInvoiceheader struct {
	db *sql.DB
}

// NewPsqlInvoiceheader return a new pointer PsqlInvoiceheader
func NewPsqlInvoiceheader(db *sql.DB) *PsqlInvoiceheader {
	return &PsqlInvoiceheader{db}
}

// Migrate implement the interface invoiceHeader.Storage
func (p *PsqlInvoiceheader) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceheader)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	log.Println("Migración de invoiceheader ejecutada correctamente")

	return nil
}
