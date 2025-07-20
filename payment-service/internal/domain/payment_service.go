package domain

import (
	"context"
	"p3-graded-challenge-1-embapge/payment-service/internal/delivery/http/dto"
)

type PaymentService interface {
	CreatePayment(ctx context.Context, req dto.CreatePaymentRequest) (*dto.CreatePaymentRequest, error)
}
