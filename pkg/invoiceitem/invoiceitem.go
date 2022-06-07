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
