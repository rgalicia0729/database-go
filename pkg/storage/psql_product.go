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

	psqlGetProductById = `
		SELECT id, name, observations, price, created_at, updated_at
		FROM products
		WHERE id = $1
	`

	psqlUpdateProduct = `
		UPDATE products
		SET name=$1, observations=$2, price=$3, updated_at=$4
		WHERE id = $5 
	`

	psqlDeleteProduct = `
		DELETE FROM products
		WHERE id = $1
	`
)

type scanner interface {
	Scan(dest ...any) error
}

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
		model, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}

		models = append(models, model)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return models, nil
}

func (p *PsqlProduct) GetById(id uint64) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlGetProductById)
	if err != nil {
		return &product.Model{}, err
	}
	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
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

func (p *PsqlProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlUpdateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		timeToNull(m.UpdatedAt),
		m.ID,
	)
	if err != nil {
		return err
	}

	if rowsAffected, err := res.RowsAffected(); err != nil {
		return err
	} else if rowsAffected == 0 {
		return product.ErrRecordNotFound
	}

	return nil
}

func (p *PsqlProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(psqlDeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func scanRowProduct(s scanner) (*product.Model, error) {
	var model product.Model

	var observationsNull sql.NullString
	var updatedAtNull sql.NullTime

	err := s.Scan(
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

	return &model, nil
}
