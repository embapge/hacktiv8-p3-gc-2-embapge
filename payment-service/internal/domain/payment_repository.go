package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *Payment) (*Payment, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*Payment, error)
	GetAll(ctx context.Context) (*[]Payment, error)
	UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}
