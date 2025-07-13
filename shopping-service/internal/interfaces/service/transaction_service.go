package service

import (
	"context"
	"p3-graded-challenge-1-embapge/shopping-service/internal/dto"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, req dto.CreateTransactionRequest) (*dto.TransactionResponse, error)
	GetAllTransactions(ctx context.Context) ([]dto.TransactionResponse, error)
	GetTransactionByID(ctx context.Context, id string) (*dto.TransactionResponse, error)
	UpdateTransaction(ctx context.Context, id string, req dto.CreateTransactionRequest) (*dto.TransactionResponse, error)
	DeleteTransaction(ctx context.Context, id string) error
}
