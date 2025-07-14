package service

import (
	"context"
	"p3-graded-challenge-1-embapge/payment-service/internal/domain"
	"p3-graded-challenge-1-embapge/payment-service/internal/dto"
	"p3-graded-challenge-1-embapge/payment-service/internal/repository"
)

type paymentService struct {
	repo repository.PaymentRepository
}

func NewPaymentService(repo repository.PaymentRepository) *paymentService {
	return &paymentService{repo: repo}
}

func (s *paymentService) CreatePayment(ctx context.Context, req dto.CreatePaymentRequest) (*dto.PaymentResponse, error) {
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
