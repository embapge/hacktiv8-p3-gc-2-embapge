package service

import (
	"context"
	"p3-graded-challenge-1-embapge/shopping-service/internal/dto"
)

type ProductService interface {
	CreateProduct(ctx context.Context, req dto.CreateProductRequest) (*dto.ProductResponse, error)
	GetAllProducts(ctx context.Context) ([]dto.ProductResponse, error)
	GetProductByID(ctx context.Context, id string) (*dto.ProductResponse, error)
	UpdateProduct(ctx context.Context, id string, req dto.UpdateProductRequest) (*dto.ProductResponse, error)
	DeleteProduct(ctx context.Context, id string) error
}
