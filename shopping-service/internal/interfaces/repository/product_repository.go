package repository

import (
	"context"
	"p3-graded-challenge-1-embapge/shopping-service/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductRepository interface {
	Create(ctx context.Context, product *domain.Product) (*domain.Product, error)
	FindAll(ctx context.Context) ([]domain.Product, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*domain.Product, error)
	Update(ctx context.Context, id primitive.ObjectID, product *domain.Product) (*domain.Product, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
}
