package repository

import (
	"context"
	"p3-graded-challenge-1-embapge/shopping-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionRepository interface {
	Create(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error)
	FindAll(ctx context.Context) ([]domain.Transaction, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*domain.Transaction, error)
	Update(ctx context.Context, id primitive.ObjectID, transaction *domain.Transaction) (*domain.Transaction, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
}
