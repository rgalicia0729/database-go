package invoiceheader

import "time"

// Model of invoiceheader
type Model struct {
	ID        uint64
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
