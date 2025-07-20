package handler

import (
	"io"
	"net/http"
	"os"
	"p3-graded-challenge-2-embapge/gateway-service/config"
	paymentpb "p3-graded-challenge-2-embapge/proto/payment"
	"strings"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GatewayHandler struct {
	GRPC *config.GRPCClients
}

func NewGatewayHandler(grpcClients *config.GRPCClients) *GatewayHandler {
	return &GatewayHandler{GRPC: grpcClients}
}

func (h *GatewayHandler) ProxyToShoppingService(c echo.Context) error {
	shoppingURL := os.Getenv("SHOPPING_SERVICE_URL")
	return proxyRequest(c, shoppingURL)
}

func (h *GatewayHandler) ProxyToAuthService(c echo.Context) error {
	authURL := os.Getenv("AUTH_SERVICE_URL")
	return proxyRequest(c, authURL)
}

func (h *GatewayHandler) ProxyToPaymentService(c echo.Context) error {
	ctx := c.Request().Context()
	method := c.Request().Method
	path := c.Request().URL.Path

	// Assume h.GRPC.PaymentClient is available and initialized

	switch method {
	case http.MethodPost:
		var req paymentpb.CreatePaymentRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}
		resp, err := h.GRPC.PaymentClient.CreatePayment(ctx, &req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusCreated, resp)
	case http.MethodGet:
		if strings.HasPrefix(path, "/payments/") && len(path) > len("/payments/") {
			id := strings.TrimPrefix(path, "/payments/")
			resp, err := h.GRPC.PaymentClient.GetByIDPayment(ctx, &paymentpb.PaymentIDRequest{Id: id})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		}
		resp, err := h.GRPC.PaymentClient.GetAllPayment(ctx, &emptypb.Empty{})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, resp)
	case http.MethodDelete:
		if strings.HasPrefix(path, "/payments/") && len(path) > len("/payments/") {
			id := strings.TrimPrefix(path, "/payments/")
			resp, err := h.GRPC.PaymentClient.DeletePayment(ctx, &paymentpb.PaymentIDRequest{Id: id})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		}
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id missing"})
	default:
		return c.NoContent(http.StatusMethodNotAllowed)
	}
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
