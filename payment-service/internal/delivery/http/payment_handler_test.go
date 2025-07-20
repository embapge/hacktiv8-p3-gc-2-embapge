package http_test

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"context"

// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// )

// type mockPaymentService struct{}

// func (m *mockPaymentService) CreatePayment(ctx context.Context, req CreatePaymentRequest) (*PaymentResponse, error) {
// 	if req.Amount <= 0 {
// 		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid amount")
// 	}
// 	return &dto.PaymentResponse{
// 		ID:     "1",
// 		Amount: req.Amount,
// 		Status: "success",
// 	}, nil
// }

// func TestCreatePayment_Success(t *testing.T) {
// 	e := echo.New()
// 	requestBody := dto.CreatePaymentRequest{Amount: 100}
// 	body, _ := json.Marshal(requestBody)
// 	req := httptest.NewRequest(http.MethodPost, "/payments", bytes.NewReader(body))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	resp := httptest.NewRecorder()
// 	c := e.NewContext(req, resp)

// 	h := handler.NewPaymentHandler(&mockPaymentService{})

// 	if assert.NoError(t, h.CreatePayment(c)) {
// 		assert.Equal(t, http.StatusCreated, resp.Code)
// 		var payment dto.PaymentResponse
// 		json.NewDecoder(resp.Body).Decode(&payment)
// 		assert.Equal(t, float64(100), payment.Amount)
// 		assert.Equal(t, "success", payment.Status)
// 	}
// }

// func TestCreatePayment_Fail_InvalidAmount(t *testing.T) {
// 	e := echo.New()
// 	requestBody := dto.CreatePaymentRequest{Amount: 0}
// 	body, _ := json.Marshal(requestBody)
// 	req := httptest.NewRequest(http.MethodPost, "/payments", bytes.NewReader(body))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	resp := httptest.NewRecorder()
// 	c := e.NewContext(req, resp)

// 	h := handler.NewPaymentHandler(&mockPaymentService{})

// 	err := h.CreatePayment(c)
// 	assert.NoError(t, err) // karena handler mengembalikan response JSON, bukan error
// 	var respBody map[string]string
// 	json.NewDecoder(resp.Body).Decode(&respBody)
// 	assert.Contains(t, respBody["message"], "required")
// }
