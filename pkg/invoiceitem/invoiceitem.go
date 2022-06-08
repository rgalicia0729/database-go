package invoiceitem

import "time"

// Model of invoiceitem
type Model struct {
	ID              uint64
	InvoiceHeaderId uint64
	ProductID       uint64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Storage interface {
	Migrate() error
}

type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{s}
}

func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
