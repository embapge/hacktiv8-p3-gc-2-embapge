package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"p3-graded-challenge-1-embapge/shopping-service/internal/dto"
	"p3-graded-challenge-1-embapge/shopping-service/internal/handler"
	mockservice "p3-graded-challenge-1-embapge/shopping-service/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct_Success(t *testing.T) {
	e := echo.New()
	requestBody := dto.CreateProductRequest{Name: "Test Product", Price: 100, Stock: 10}
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)

	mockSvc := &mockservice.MockProductService{
		CreateProductFunc: func(ctx context.Context, req dto.CreateProductRequest) (*dto.ProductResponse, error) {
			return &dto.ProductResponse{
				ID:    "1",
				Name:  req.Name,
				Price: req.Price,
				Stock: req.Stock,
			}, nil
		},
	}
	h := handler.NewProductHandler(mockSvc)

	if assert.NoError(t, h.CreateProduct(c)) {
		assert.Equal(t, http.StatusCreated, resp.Code)
		var product dto.ProductResponse
		json.NewDecoder(resp.Body).Decode(&product)
		assert.Equal(t, "Test Product", product.Name)
		assert.Equal(t, float64(100), product.Price)
		assert.Equal(t, 10, product.Stock)
	}
}

func TestCreateProduct_Fail_InvalidName(t *testing.T) {
	e := echo.New()
	requestBody := dto.CreateProductRequest{Name: "", Price: 100, Stock: 10}
	body, _ := json.Marshal(requestBody)
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)

	// Arrange: mock service returns error
	mockSvc := &mockservice.MockProductService{
		CreateProductFunc: func(ctx context.Context, req dto.CreateProductRequest) (*dto.ProductResponse, error) {
			return nil, assert.AnError
		},
	}
	h := handler.NewProductHandler(mockSvc)

	err := h.CreateProduct(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
