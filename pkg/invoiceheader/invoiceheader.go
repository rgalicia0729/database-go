package invoiceheader

import "time"

// Model of invoiceheader
type Model struct {
	ID        uint64
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Storage interface {
	Migrate() error
}

// Service of invoiceheader
type Service struct {
	storage Storage
}

// NewService return a pointer with Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used for migrate invoiceheader
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
