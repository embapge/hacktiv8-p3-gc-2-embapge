package http

import (
	"net/http"
	"p3-graded-challenge-2-embapge/payment-service/internal/app"
	"p3-graded-challenge-2-embapge/payment-service/internal/delivery/http/dto"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	app       *app.PaymentApp
	validator *validator.Validate
}

func NewPaymentHandler(app *app.PaymentApp) *PaymentHandler {
	return &PaymentHandler{
		app:       app,
		validator: validator.New(),
	}
}

func (h *PaymentHandler) CreatePaymentHandler(c echo.Context) error {
	var req dto.CreatePaymentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}
	if err := h.validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	payment, err := h.app.CreatePayment(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create payment"})
	}
	return c.JSON(http.StatusCreated, payment)
}
