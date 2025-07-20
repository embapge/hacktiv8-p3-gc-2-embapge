package domain

import (
	"context"
	"p3-graded-challenge-2-embapge/payment-service/internal/delivery/http/dto"
)

type PaymentService interface {
	CreatePayment(ctx context.Context, req dto.CreatePaymentRequest) (*dto.PaymentResponse, error)
	FindByIDPayment(ctx context.Context, id string) (*Payment, error)
	GetAllPayment(ctx context.Context) (*[]Payment, error)
	Delete(ctx context.Context, id string) error
}
