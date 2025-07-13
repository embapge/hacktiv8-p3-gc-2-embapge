package router

import (
	"p3-graded-challenge-1-embapge/shopping-service/internal/interfaces/handler"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(e *echo.Echo, productHandler handler.ProductHandler, transactionHandler handler.TransactionHandler) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	g := e.Group("/products")
	g.POST("", productHandler.CreateProduct)
	g.GET("", productHandler.GetAllProducts)
	g.GET("/:id", productHandler.GetProductByID)
	g.PUT("/:id", productHandler.UpdateProduct)
	g.DELETE("/:id", productHandler.DeleteProduct)

	t := e.Group("/transactions")
	t.POST("", transactionHandler.CreateTransaction)
	t.GET("", transactionHandler.GetAllTransactions)
	t.GET("/:id", transactionHandler.GetTransactionByID)
	t.PUT("/:id", transactionHandler.UpdateTransaction)
	t.DELETE("/:id", transactionHandler.DeleteTransaction)
}
