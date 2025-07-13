package handler

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type GatewayHandler struct{}

func NewGatewayHandler() *GatewayHandler {
	return &GatewayHandler{}
}

func (h *GatewayHandler) ProxyToShoppingService(c echo.Context) error {
	shoppingURL := os.Getenv("SHOPPING_SERVICE_URL")
	resp, err := http.Get(shoppingURL + c.Request().RequestURI)
	if err != nil {
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "Failed to connect to shopping service"})
	}
	defer resp.Body.Close()
	return proxyResponse(c, resp)
}

func (h *GatewayHandler) ProxyToPaymentService(c echo.Context) error {
	paymentURL := os.Getenv("PAYMENT_SERVICE_URL")
	resp, err := http.Get(paymentURL + c.Request().RequestURI)
	if err != nil {
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "Failed to connect to payment service"})
	}
	defer resp.Body.Close()
	return proxyResponse(c, resp)
}

func proxyResponse(c echo.Context, resp *http.Response) error {
	for k, v := range resp.Header {
		for _, vv := range v {
			c.Response().Header().Add(k, vv)
		}
	}
	c.Response().WriteHeader(resp.StatusCode)
	_, err := io.Copy(c.Response(), resp.Body)
	return err
}
