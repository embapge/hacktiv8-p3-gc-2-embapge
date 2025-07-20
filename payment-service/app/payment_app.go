package app

import (
	"context"
	"p3-graded-challenge-1-embapge/payment-service/internal/delivery/http/dto"
	"p3-graded-challenge-1-embapge/payment-service/internal/domain"
)

type PaymentApp struct {
	repo domain.PaymentRepository
}

func NewPaymentApp(repo domain.PaymentRepository) *PaymentApp {
	return &PaymentApp{repo: repo}
}

func (s *PaymentApp) CreatePayment(ctx context.Context, req dto.CreatePaymentRequest) (*dto.PaymentResponse, error) {
	payment := &domain.Payment{
		Amount: req.Amount,
	}
	created, err := s.repo.Create(ctx, payment)
	if err != nil {
		return nil, err
	}
	return &dto.PaymentResponse{
		ID:        created.ID.Hex(),
		Amount:    created.Amount,
		Status:    created.Status,
		CreatedAt: created.CreatedAt,
		UpdatedAt: created.UpdatedAt,
	}, nil
}
