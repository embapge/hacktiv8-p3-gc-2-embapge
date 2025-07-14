package main

import (
	"log"
	"os"

	_ "p3-graded-challenge-1-embapge/gateway-service/docs"
	"p3-graded-challenge-1-embapge/gateway-service/handler"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Gateway API
// @version 1.0
// @description API Gateway for Shopping and Payment Services
// @host 34.101.156.80:8000
// @BasePath /

// ===== Swagger Documentation for Shopping Service =====

// @Summary Create a new product
// @Description Create a new product with the input payload
// @Tags products
// @Accept json
// @Produce json
// @Param product body CreateProductRequest true "Create Product"
// @Success 201 {object} ProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /products [post]
func _SwaggerCreateProduct() {}

// @Summary Get all products
// @Description Get all products
// @Tags products
// @Produce json
// @Success 200 {array} ProductResponse
// @Failure 500 {object} ErrorResponse
// @Router /products [get]
func _SwaggerGetAllProducts() {}

// @Summary Get product by ID
// @Description Get product by ID
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} ProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /products/{id} [get]
func _SwaggerGetProductByID() {}

// @Summary Update a product
// @Description Update a product with the input payload
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body UpdateProductRequest true "Update Product"
// @Success 200 {object} ProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /products/{id} [put]
func _SwaggerUpdateProduct() {}

// @Summary Delete a product
// @Description Delete a product by ID
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /products/{id} [delete]
func _SwaggerDeleteProduct() {}

// @Summary Create a new transaction
// @Description Create a new transaction and payment
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction body CreateTransactionRequest true "Create Transaction"
// @Success 201 {object} TransactionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transactions [post]
func _SwaggerCreateTransaction() {}

// @Summary Get all transactions
// @Description Get all transactions
// @Tags transactions
// @Produce json
// @Success 200 {array} TransactionResponse
// @Failure 500 {object} ErrorResponse
// @Router /transactions [get]
func _SwaggerGetAllTransactions() {}

// @Summary Get transaction by ID
// @Description Get transaction by ID
// @Tags transactions
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} TransactionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transactions/{id} [get]
func _SwaggerGetTransactionByID() {}

// @Summary Update a transaction
// @Description Update a transaction with the input payload
// @Tags transactions
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Param transaction body CreateTransactionRequest true "Update Transaction"
// @Success 200 {object} TransactionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transactions/{id} [put]
func _SwaggerUpdateTransaction() {}

// @Summary Delete a transaction
// @Description Delete a transaction by ID
// @Tags transactions
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transactions/{id} [delete]
func _SwaggerDeleteTransaction() {}

// ===== Swagger Documentation for Payment Service =====

// @Summary Create payment
// @Description Create payment via payment-service
// @Tags payments
// @Accept json
// @Produce json
// @Param payment body CreatePaymentRequest true "Create Payment"
// @Success 201 {object} PaymentResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /payments [post]
func _SwaggerCreatePayment() {}

// ===== DTOs for Swagger =====
type CreateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}
type UpdateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}
type ProductResponse struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
type CreateTransactionRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
type TransactionResponse struct {
	ID        string  `json:"id"`
	ProductID string  `json:"product_id"`
	PaymentID string  `json:"payment_id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
type CreatePaymentRequest struct {
	Amount float64 `json:"amount"`
}
type PaymentResponse struct {
	ID        string  `json:"id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
type ErrorResponse struct {
	Message string `json:"message"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	h := handler.NewGatewayHandler()

	// Proxy all /products* and /transactions* to shopping-service
	e.Any("/products*", h.ProxyToShoppingService)
	e.Any("/transactions*", h.ProxyToShoppingService)

	// Proxy all /payments* to payment-service
	e.Any("/payments*", h.ProxyToPaymentService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Println("Gateway running on port", port)
	e.Logger.Fatal(e.Start(":" + port))
}
