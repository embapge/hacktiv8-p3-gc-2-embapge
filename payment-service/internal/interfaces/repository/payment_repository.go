package repository

import (
	"context"
	"p3-graded-challenge-1-embapge/payment-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *domain.Payment) (*domain.Payment, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*domain.Payment, error)
	UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error
}
