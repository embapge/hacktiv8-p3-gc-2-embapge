package service

import (
	"context"
	"errors"
	"time"

	"p3-graded-challenge-1-embapge/shopping-service/internal/domain"
	"p3-graded-challenge-1-embapge/shopping-service/internal/dto"
	"p3-graded-challenge-1-embapge/shopping-service/internal/interfaces/repository"
	"p3-graded-challenge-1-embapge/shopping-service/internal/interfaces/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) service.ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) CreateProduct(ctx context.Context, req dto.CreateProductRequest) (*dto.ProductResponse, error) {
	product := &domain.Product{
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
	}

	createdProduct, err := s.repo.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:        createdProduct.ID.Hex(),
		Name:      createdProduct.Name,
		Price:     createdProduct.Price,
		Stock:     createdProduct.Stock,
		CreatedAt: createdProduct.CreatedAt,
	}, nil
}

func (s *productService) GetAllProducts(ctx context.Context) ([]dto.ProductResponse, error) {
	products, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var productResponses []dto.ProductResponse
	for _, p := range products {
		productResponses = append(productResponses, dto.ProductResponse{
			ID:        p.ID.Hex(),
			Name:      p.Name,
			Price:     p.Price,
			Stock:     p.Stock,
			CreatedAt: p.CreatedAt,
		})
	}
	return productResponses, nil
}

func (s *productService) GetProductByID(ctx context.Context, id string) (*dto.ProductResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid product id")
	}

	product, err := s.repo.FindByID(ctx, objID)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	return &dto.ProductResponse{
		ID:        product.ID.Hex(),
		Name:      product.Name,
		Price:     product.Price,
		Stock:     product.Stock,
		CreatedAt: product.CreatedAt,
	}, nil
}

func (s *productService) UpdateProduct(ctx context.Context, id string, req dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid product id")
	}

	product, err := s.repo.FindByID(ctx, objID)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Price > 0 {
		product.Price = req.Price
	}
	if req.Stock >= 0 {
		product.Stock = req.Stock
	}
	product.UpdatedAt = time.Now()

	updatedProduct, err := s.repo.Update(ctx, objID, product)
	if err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:        updatedProduct.ID.Hex(),
		Name:      updatedProduct.Name,
		Price:     updatedProduct.Price,
		Stock:     updatedProduct.Stock,
		CreatedAt: updatedProduct.CreatedAt,
		UpdatedAt: updatedProduct.UpdatedAt,
	}, nil
}

func (s *productService) DeleteProduct(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid product id")
	}

	product, err := s.repo.FindByID(ctx, objID)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("product not found")
	}

	return s.repo.Delete(ctx, objID)
}
