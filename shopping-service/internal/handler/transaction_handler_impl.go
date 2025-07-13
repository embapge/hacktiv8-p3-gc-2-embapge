package handler

import (
	"net/http"
	"p3-graded-challenge-1-embapge/shopping-service/internal/dto"
	"p3-graded-challenge-1-embapge/shopping-service/internal/interfaces/handler"
	"p3-graded-challenge-1-embapge/shopping-service/internal/interfaces/service"
	"p3-graded-challenge-1-embapge/shopping-service/pkg/utility"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type transactionHandler struct {
	service   service.TransactionService
	validator *validator.Validate
}

func NewTransactionHandler(service service.TransactionService) handler.TransactionHandler {
	return &transactionHandler{
		service:   service,
		validator: validator.New(),
	}
}

// CreateTransaction godoc
// @Summary Create a new transaction
// @Description Create a new transaction and payment
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param transaction body dto.CreateTransactionRequest true "Create Transaction"
// @Success 201 {object} dto.TransactionResponse
// @Failure 400 {object} utility.ErrorResponse
// @Failure 500 {object} utility.ErrorResponse
// @Router /transactions [post]
func (h *transactionHandler) CreateTransaction(c echo.Context) error {
	var req dto.CreateTransactionRequest
	if err := c.Bind(&req); err != nil {
		return utility.BadRequestError(c, "Invalid request body")
	}
	if err := h.validator.Struct(req); err != nil {
		return utility.BadRequestError(c, err.Error())
	}
	trx, err := h.service.CreateTransaction(c.Request().Context(), req)
	if err != nil {
		return utility.InternalServerError(c, err.Error())
	}
	return c.JSON(http.StatusCreated, trx)
}

func (h *transactionHandler) GetAllTransactions(c echo.Context) error {
	trxs, err := h.service.GetAllTransactions(c.Request().Context())
	if err != nil {
		return utility.InternalServerError(c, err.Error())
	}
	return c.JSON(http.StatusOK, trxs)
}

func (h *transactionHandler) GetTransactionByID(c echo.Context) error {
	id := c.Param("id")
	trx, err := h.service.GetTransactionByID(c.Request().Context(), id)
	if err != nil {
		if err.Error() == "invalid transaction id" {
			return utility.BadRequestError(c, err.Error())
		}
		if err.Error() == "transaction not found" {
			return utility.NotFoundError(c, err.Error())
		}
		return utility.InternalServerError(c, err.Error())
	}
	return c.JSON(http.StatusOK, trx)
}

func (h *transactionHandler) UpdateTransaction(c echo.Context) error {
	id := c.Param("id")
	var req dto.CreateTransactionRequest
	if err := c.Bind(&req); err != nil {
		return utility.BadRequestError(c, "Invalid request body")
	}
	if err := h.validator.Struct(req); err != nil {
		return utility.BadRequestError(c, err.Error())
	}
	trx, err := h.service.UpdateTransaction(c.Request().Context(), id, req)
	if err != nil {
		if err.Error() == "invalid transaction id" {
			return utility.BadRequestError(c, err.Error())
		}
		if err.Error() == "transaction not found" {
			return utility.NotFoundError(c, err.Error())
		}
		return utility.InternalServerError(c, err.Error())
	}
	return c.JSON(http.StatusOK, trx)
}

func (h *transactionHandler) DeleteTransaction(c echo.Context) error {
	id := c.Param("id")
	err := h.service.DeleteTransaction(c.Request().Context(), id)
	if err != nil {
		if err.Error() == "invalid transaction id" {
			return utility.BadRequestError(c, err.Error())
		}
		if err.Error() == "transaction not found" {
			return utility.NotFoundError(c, err.Error())
		}
		return utility.InternalServerError(c, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
