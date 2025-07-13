package handler

import "github.com/labstack/echo/v4"

type TransactionHandler interface {
	CreateTransaction(c echo.Context) error
	GetAllTransactions(c echo.Context) error
	GetTransactionByID(c echo.Context) error
	UpdateTransaction(c echo.Context) error
	DeleteTransaction(c echo.Context) error
}
