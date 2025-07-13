package service

import (
	"context"
	"p3-graded-challenge-1-embapge/payment-service/internal/dto"
)

type PaymentService interface {
	CreatePayment(ctx context.Context, req dto.CreatePaymentRequest) (*dto.PaymentResponse, error)
}
