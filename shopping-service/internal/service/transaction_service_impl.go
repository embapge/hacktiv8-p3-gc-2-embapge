package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"p3-graded-challenge-1-embapge/shopping-service/internal/domain"
	"p3-graded-challenge-1-embapge/shopping-service/internal/dto"
	"p3-graded-challenge-1-embapge/shopping-service/internal/interfaces/repository"
	"p3-graded-challenge-1-embapge/shopping-service/internal/interfaces/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type transactionService struct {
	txRepo      repository.TransactionRepository
	productRepo repository.ProductRepository
}

func NewTransactionService(txRepo repository.TransactionRepository, productRepo repository.ProductRepository) service.TransactionService {
	return &transactionService{
		txRepo:      txRepo,
		productRepo: productRepo,
	}
}

func (s *transactionService) CreateTransaction(ctx context.Context, req dto.CreateTransactionRequest) (*dto.TransactionResponse, error) {
	// Validasi produk
	productID, err := primitive.ObjectIDFromHex(req.ProductID)
	if err != nil {
		return nil, errors.New("invalid product id")
	}
	product, err := s.productRepo.FindByID(ctx, productID)
	if err != nil || product == nil {
		return nil, errors.New("product not found")
	}
	if product.Stock < req.Quantity {
		return nil, errors.New("insufficient stock")
	}
	total := product.Price * float64(req.Quantity)

	// Call Payment Service
	paymentReq := map[string]interface{}{
		"amount": total,
	}
	paymentBody, _ := json.Marshal(paymentReq)
	paymentServiceURL := os.Getenv("PAYMENT_SERVICE_URL")
	if paymentServiceURL == "" {
		log.Fatal("PAYMENT_SERVICE_URL is not set")
	}
	resp, err := http.Post(paymentServiceURL, "application/json", bytes.NewBuffer(paymentBody))
	if err != nil || resp.StatusCode != http.StatusCreated {
		return nil, errors.New("failed to create payment")
	}
	var paymentResp struct {
		ID        string    `json:"id"`
		Amount    float64   `json:"amount"`
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&paymentResp); err != nil {
		return nil, errors.New("invalid payment response")
	}
	if paymentResp.Status != "pending" {
		return nil, errors.New("payment not pending")
	}

	// Kurangi stok produk
	product.Stock -= req.Quantity
	product.UpdatedAt = time.Now()
	_, err = s.productRepo.Update(ctx, productID, product)
	if err != nil {
		return nil, errors.New("failed to update product stock")
	}

	paymentObjID, err := primitive.ObjectIDFromHex(paymentResp.ID)
	if err != nil {
		return nil, errors.New("invalid payment id format")
	}

	tx := &domain.Transaction{
		ProductID: productID,
		PaymentID: paymentObjID,
		Quantity:  req.Quantity,
		Total:     total,
		Status:    "success", // default success, akan diubah jika payment gagal
	}
	createdTx, err := s.txRepo.Create(ctx, tx)
	if err != nil {
		return nil, err
	}

	return &dto.TransactionResponse{
		ID:        createdTx.ID.Hex(),
		ProductID: createdTx.ProductID.Hex(),
		PaymentID: createdTx.PaymentID.Hex(),
		Quantity:  createdTx.Quantity,
		Total:     createdTx.Total,
		Status:    createdTx.Status,
		CreatedAt: createdTx.CreatedAt,
		UpdatedAt: createdTx.UpdatedAt,
	}, nil
}

func (s *transactionService) GetAllTransactions(ctx context.Context) ([]dto.TransactionResponse, error) {
	transactions, err := s.txRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var responses []dto.TransactionResponse
	for _, t := range transactions {
		responses = append(responses, dto.TransactionResponse{
			ID:        t.ID.Hex(),
			ProductID: t.ProductID.Hex(),
			PaymentID: t.PaymentID.Hex(),
			Quantity:  t.Quantity,
			Total:     t.Total,
			Status:    t.Status,
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
		})
	}
	return responses, nil
}

func (s *transactionService) GetTransactionByID(ctx context.Context, id string) (*dto.TransactionResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid transaction id")
	}
	tx, err := s.txRepo.FindByID(ctx, objID)
	if err != nil {
		return nil, err
	}
	if tx == nil {
		return nil, errors.New("transaction not found")
	}
	return &dto.TransactionResponse{
		ID:        tx.ID.Hex(),
		ProductID: tx.ProductID.Hex(),
		PaymentID: tx.PaymentID.Hex(),
		Quantity:  tx.Quantity,
		Total:     tx.Total,
		Status:    tx.Status,
		CreatedAt: tx.CreatedAt,
		UpdatedAt: tx.UpdatedAt,
	}, nil
}

func (s *transactionService) UpdateTransaction(ctx context.Context, id string, req dto.CreateTransactionRequest) (*dto.TransactionResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid transaction id")
	}
	tx, err := s.txRepo.FindByID(ctx, objID)
	if err != nil {
		return nil, err
	}
	if tx == nil {
		return nil, errors.New("transaction not found")
	}
	if req.Quantity > 0 {
		tx.Quantity = req.Quantity
	}
	tx.UpdatedAt = time.Now()
	updatedTx, err := s.txRepo.Update(ctx, objID, tx)
	if err != nil {
		return nil, err
	}
	return &dto.TransactionResponse{
		ID:        updatedTx.ID.Hex(),
		ProductID: updatedTx.ProductID.Hex(),
		PaymentID: updatedTx.PaymentID.Hex(),
		Quantity:  updatedTx.Quantity,
		Total:     updatedTx.Total,
		Status:    updatedTx.Status,
		CreatedAt: updatedTx.CreatedAt,
		UpdatedAt: updatedTx.UpdatedAt,
	}, nil
}

func (s *transactionService) DeleteTransaction(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid transaction id")
	}
	tx, err := s.txRepo.FindByID(ctx, objID)
	if err != nil {
		return err
	}
	if tx == nil {
		return errors.New("transaction not found")
	}
	return s.txRepo.Delete(ctx, objID)
}
