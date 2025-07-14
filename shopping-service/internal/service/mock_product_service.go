package service

import (
	"context"
	"p3-graded-challenge-1-embapge/shopping-service/internal/dto"
)

type MockProductService struct {
	CreateProductFunc  func(ctx context.Context, req dto.CreateProductRequest) (*dto.ProductResponse, error)
	DeleteProductFunc  func(ctx context.Context, id string) error
	GetAllProductsFunc func(ctx context.Context) ([]dto.ProductResponse, error)
	GetProductByIDFunc func(ctx context.Context, id string) (*dto.ProductResponse, error)
	UpdateProductFunc  func(ctx context.Context, id string, req dto.UpdateProductRequest) (*dto.ProductResponse, error)
}

func (m *MockProductService) CreateProduct(ctx context.Context, req dto.CreateProductRequest) (*dto.ProductResponse, error) {
	return m.CreateProductFunc(ctx, req)
}

func (m *MockProductService) DeleteProduct(ctx context.Context, id string) error {
	return m.DeleteProductFunc(ctx, id)
}

func (m *MockProductService) GetAllProducts(ctx context.Context) ([]dto.ProductResponse, error) {
	return m.GetAllProductsFunc(ctx)
}

func (m *MockProductService) GetProductByID(ctx context.Context, id string) (*dto.ProductResponse, error) {
	return m.GetProductByIDFunc(ctx, id)
}

// UpdateProduct implements ProductService.UpdateProduct
func (m *MockProductService) UpdateProduct(ctx context.Context, id string, req dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	return m.UpdateProductFunc(ctx, id, req)
}
