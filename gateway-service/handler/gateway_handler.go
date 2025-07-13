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
	return proxyRequest(c, shoppingURL)
}

func (h *GatewayHandler) ProxyToPaymentService(c echo.Context) error {
	paymentURL := os.Getenv("PAYMENT_SERVICE_URL")
	return proxyRequest(c, paymentURL)
}

func proxyRequest(c echo.Context, targetURL string) error {
	req, err := http.NewRequest(c.Request().Method, targetURL+c.Request().RequestURI, c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create proxy request"})
	}

	req.Header = c.Request().Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "Failed to connect to the downstream service"})
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
