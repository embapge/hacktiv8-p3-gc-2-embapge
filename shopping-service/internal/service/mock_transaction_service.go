package service

import (
	"context"
	"p3-graded-challenge-1-embapge/shopping-service/internal/dto"
)

// MockTransactionService implements the TransactionService interface for testing.
type MockTransactionService struct {
	CreateTransactionFunc  func(ctx context.Context, req dto.CreateTransactionRequest) (*dto.TransactionResponse, error)
	GetAllTransactionsFunc func(ctx context.Context) ([]dto.TransactionResponse, error)
	GetTransactionByIDFunc func(ctx context.Context, id string) (*dto.TransactionResponse, error)
	UpdateTransactionFunc  func(ctx context.Context, id string, req dto.CreateTransactionRequest) (*dto.TransactionResponse, error)
	DeleteTransactionFunc  func(ctx context.Context, id string) error
}

// CreateTransaction calls the mocked function.
func (m *MockTransactionService) CreateTransaction(ctx context.Context, req dto.CreateTransactionRequest) (*dto.TransactionResponse, error) {
	return m.CreateTransactionFunc(ctx, req)
}

// GetAllTransactions calls the mocked function.
func (m *MockTransactionService) GetAllTransactions(ctx context.Context) ([]dto.TransactionResponse, error) {
	return m.GetAllTransactionsFunc(ctx)
}

// GetTransactionByID calls the mocked function.
func (m *MockTransactionService) GetTransactionByID(ctx context.Context, id string) (*dto.TransactionResponse, error) {
	return m.GetTransactionByIDFunc(ctx, id)
}

// UpdateTransaction calls the mocked function.
func (m *MockTransactionService) UpdateTransaction(ctx context.Context, id string, req dto.CreateTransactionRequest) (*dto.TransactionResponse, error) {
	return m.UpdateTransactionFunc(ctx, id, req)
}

// DeleteTransaction calls the mocked function.
func (m *MockTransactionService) DeleteTransaction(ctx context.Context, id string) error {
	return m.DeleteTransactionFunc(ctx, id)
}
