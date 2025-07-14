package handler

import (
	"net/http"
	"p3-graded-challenge-1-embapge/payment-service/internal/dto"
	"p3-graded-challenge-1-embapge/payment-service/internal/interfaces/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type paymentHandler struct {
	service   service.PaymentService
	validator *validator.Validate
}

func NewPaymentHandler(service service.PaymentService) *paymentHandler {
	return &paymentHandler{
		service:   service,
		validator: validator.New(),
	}
}

func (h *paymentHandler) CreatePayment(c echo.Context) error {
	var req dto.CreatePaymentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}
	if err := h.validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	payment, err := h.service.CreatePayment(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create payment"})
	}
	return c.JSON(http.StatusCreated, payment)
}
