package app

import (
	"context"
	"errors"
	"p3-graded-challenge-2-embapge/payment-service/internal/delivery/http/dto"
	"p3-graded-challenge-2-embapge/payment-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (s *PaymentApp) FindByIDPayment(ctx context.Context, id string) (*dto.PaymentResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid payment id")
	}

	payment, err := s.repo.FindByID(ctx, objID)
	if err != nil {
		return nil, err
	}

	return &dto.PaymentResponse{
		ID:        payment.ID.Hex(),
		Amount:    payment.Amount,
		Status:    payment.Status,
		CreatedAt: payment.CreatedAt,
		UpdatedAt: payment.UpdatedAt,
	}, nil
}

func (s *PaymentApp) GetAllPayment(ctx context.Context) (*[]domain.Payment, error) {
	payments, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (s *PaymentApp) DeletePayment(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid payment id")
	}

	err = s.repo.Delete(ctx, objID)
	if err != nil {
		return err
	}

	return nil
}
