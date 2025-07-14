package router

import (
	"p3-graded-challenge-1-embapge/payment-service/internal/interfaces/handler"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, paymentHandler handler.PaymentHandler) {
	e.POST("/payments", paymentHandler.CreatePayment)
}
