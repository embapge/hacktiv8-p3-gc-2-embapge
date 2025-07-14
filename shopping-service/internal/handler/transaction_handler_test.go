package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"p3-graded-challenge-1-embapge/shopping-service/internal/dto"
	"p3-graded-challenge-1-embapge/shopping-service/internal/handler"
	mockservice "p3-graded-challenge-1-embapge/shopping-service/internal/service"
)

func TestCreateTransaction_Success(t *testing.T) {
	e := echo.New()
	requestBody := dto.CreateTransactionRequest{ProductID: "prod1", Quantity: 2}
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)

	mockSvc := &mockservice.MockTransactionService{
		CreateTransactionFunc: func(ctx context.Context, req dto.CreateTransactionRequest) (*dto.TransactionResponse, error) {
			return &dto.TransactionResponse{
				ID:        "trx1",
				ProductID: req.ProductID,
				PaymentID: "pay1",
				Quantity:  req.Quantity,
				Total:     200.0,
				Status:    "created",
			}, nil
		},
	}
	h := handler.NewTransactionHandler(mockSvc)

	err := h.CreateTransaction(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.Code)

	var trx dto.TransactionResponse
	json.NewDecoder(resp.Body).Decode(&trx)
	assert.Equal(t, "trx1", trx.ID)
	assert.Equal(t, "prod1", trx.ProductID)
	assert.Equal(t, "pay1", trx.PaymentID)
	assert.Equal(t, 2, trx.Quantity)
	assert.Equal(t, 200.0, trx.Total)
	assert.Equal(t, "created", trx.Status)
}

func TestCreateTransaction_ServiceError(t *testing.T) {
	e := echo.New()
	requestBody := dto.CreateTransactionRequest{ProductID: "prod1", Quantity: 2}
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)

	mockSvc := &mockservice.MockTransactionService{
		CreateTransactionFunc: func(ctx context.Context, req dto.CreateTransactionRequest) (*dto.TransactionResponse, error) {
			return nil, assert.AnError
		},
	}
	h := handler.NewTransactionHandler(mockSvc)

	err := h.CreateTransaction(c)
	// handler returns error via JSON, but method should not panic
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}
