package dto

import "time"

type CreatePaymentRequest struct {
	Amount float64 `json:"amount" validate:"required,gt=0"`
}

type PaymentResponse struct {
	ID        string    `json:"id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
