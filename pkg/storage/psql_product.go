package storage

import (
	"database/sql"
	"log"

	"github.com/rgalicia0729/database-go/pkg/product"
)

const (
	psqlMigrateProduct = `
		CREATE TABLE IF NOT EXISTS products(
			id SERIAL NOT NULL,
			name CHARACTER VARYING(25) NOT NULL,
			observations CHARACTER VARYING(100),
			price DECIMAL(12,2) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP,
			CONSTRAINT products_id_pk PRIMARY KEY (id)
		)
	`

	psqlCreateProduct = `
		INSERT INTO products(name, observations, price)
		VALUES($1, $2, $3)
		RETURNING id
	`

	psqlGetAllProduct = `
		SELECT id, name, observations, price, created_at, updated_at
		FROM products
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

func (p *PsqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var models product.Models
	for rows.Next() {
		var model product.Model

		var observationsNull sql.NullString
		var updatedAtNull sql.NullTime

		err := rows.Scan(
			&model.ID,
			&model.Name,
			&observationsNull,
			&model.Price,
			&model.CreatedAt,
			&updatedAtNull,
		)
		if err != nil {
			return nil, err
		}

		model.Observations = observationsNull.String
		model.UpdatedAt = updatedAtNull.Time

		models = append(models, &model)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return models, nil
}

func (p *PsqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
	).Scan(&m.ID)
	if err != nil {
		return err
	}

	log.Println("Se creo el producto correctamente")

	return nil
}
