package handler

import "github.com/labstack/echo/v4"

type ProductHandler interface {
	CreateProduct(c echo.Context) error
	GetAllProducts(c echo.Context) error
	GetProductByID(c echo.Context) error
	UpdateProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
}
