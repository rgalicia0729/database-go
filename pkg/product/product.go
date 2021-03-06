package product

import (
	"errors"
	"fmt"
	"time"
)

// Model of product
type Model struct {
	ID           uint64
	Name         string
	Observations string
	Price        float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (m *Model) String() string {
	return fmt.Sprintf(
		"%02d | %-20s | %-30s | %10f | %10s | %10s\n",
		m.ID,
		m.Name,
		m.Observations,
		m.Price,
		m.CreatedAt.Format("2006-01-02"),
		m.UpdatedAt.Format("2006-01-02"),
	)
}

// Models slice of Model
type Models []*Model

// Storage
type Storage interface {
	Migrate() error
	Create(*Model) error
	GetAll() (Models, error)
	GetById(uint64) (*Model, error)
	Update(*Model) error
	Delete(uint) error
}

var (
	ErrIdNotFound     = errors.New("error: id not found")
	ErrRecordNotFound = errors.New("error: record not found")
)

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

func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}

func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

func (s *Service) GetById(id uint64) (*Model, error) {
	return s.storage.GetById(id)
}

func (s *Service) Update(m *Model) error {
	if m.ID == 0 {
		return ErrIdNotFound
	}

	m.UpdatedAt = time.Now()

	return s.storage.Update(m)
}

func (s *Service) Delete(id uint) error {
	return s.storage.Delete(id)
}
