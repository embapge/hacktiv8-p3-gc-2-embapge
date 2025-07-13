package dto

import "time"

type CreateTransactionRequest struct {
	ProductID string `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,gte=1"`
}

type TransactionResponse struct {
	ID        string    `json:"id"`
	ProductID string    `json:"product_id"`
	PaymentID string    `json:"payment_id"`
	Quantity  int       `json:"quantity"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
