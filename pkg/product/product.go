package product

import "time"

// Model of product
type Model struct {
	ID           uint64
	Name         string
	Observations string
	Price        float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Models slice of Model
type Models []*Model

// Storage
type Storage interface {
	Migrate() error
}

// Service of product
type Service struct {
	storage Storage
}

// NewService return a pointer of service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used for migrate product
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
