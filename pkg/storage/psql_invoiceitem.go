package storage

import (
	"database/sql"
	"log"
)

const (
	psqlMigrateInvoiceitem = `
		CREATE TABLE IF NOT EXISTS invoice_items(
			id SERIAL NOT NULL,
			invoice_header_id INT NOT NULL,
			product_id INT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP,
			CONSTRAINT invoice_item_id_pk PRIMARY KEY (id),
			CONSTRAINT invoice_item_invoice_header_id_fk FOREIGN KEY 
				(invoice_header_id) REFERENCES invoice_headers(id)
				ON UPDATE RESTRICT
				ON DELETE RESTRICT,
			CONSTRAINT invoice_item_product_id_fk FOREIGN KEY
				(product_id) REFERENCES products(id)
				ON UPDATE RESTRICT
				ON DELETE RESTRICT
		)
	`
)

type psqlInvoiceitem struct {
	db *sql.DB
}

func NewPsqlInvoiceitem(db *sql.DB) *psqlInvoiceitem {
	return &psqlInvoiceitem{db}
}

func (p *psqlInvoiceitem) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceitem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	log.Println("Migración de invoiceitem ejecutada correctamente")

	return nil

}
