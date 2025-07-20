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

type productHandler struct {
	service   service.ProductService
	validator *validator.Validate
}

func NewProductHandler(service service.ProductService) handler.ProductHandler {
	return &productHandler{
		service:   service,
		validator: validator.New(),
	}
}

func (h *productHandler) CreateProduct(c echo.Context) error {
	var req dto.CreateProductRequest
	if err := c.Bind(&req); err != nil {
		return utility.BadRequestError(c, "Invalid request body")
	}

	if err := h.validator.Struct(req); err != nil {
		return utility.BadRequestError(c, err.Error())
	}

	product, err := h.service.CreateProduct(c.Request().Context(), req)
	if err != nil {
		return utility.InternalServerError(c, "Failed to create product")
	}

	return c.JSON(http.StatusCreated, product)
}

func (h *productHandler) GetAllProducts(c echo.Context) error {
	products, err := h.service.GetAllProducts(c.Request().Context())
	if err != nil {
		return utility.InternalServerError(c, "Failed to retrieve products")
	}
	return c.JSON(http.StatusOK, products)
}
func (h *productHandler) GetProductByID(c echo.Context) error {
	id := c.Param("id")
	product, err := h.service.GetProductByID(c.Request().Context(), id)
	if err != nil {
		if err.Error() == "invalid product id" {
			return utility.BadRequestError(c, err.Error())
		}
		if err.Error() == "product not found" {
			return utility.NotFoundError(c, err.Error())
		}
		return utility.InternalServerError(c, "Failed to retrieve product")
	}
	return c.JSON(http.StatusOK, product)
}

func (h *productHandler) UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	var req dto.UpdateProductRequest
	if err := c.Bind(&req); err != nil {
		return utility.BadRequestError(c, "Invalid request body")
	}

	if err := h.validator.Struct(req); err != nil {
		return utility.BadRequestError(c, err.Error())
	}

	product, err := h.service.UpdateProduct(c.Request().Context(), id, req)
	if err != nil {
		if err.Error() == "invalid product id" {
			return utility.BadRequestError(c, err.Error())
		}
		if err.Error() == "product not found" {
			return utility.NotFoundError(c, err.Error())
		}
		return utility.InternalServerError(c, "Failed to update product")
	}
	return c.JSON(http.StatusOK, product)
}

func (h *productHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	err := h.service.DeleteProduct(c.Request().Context(), id)
	if err != nil {
		if err.Error() == "invalid product id" {
			return utility.BadRequestError(c, err.Error())
		}
		if err.Error() == "product not found" {
			return utility.NotFoundError(c, err.Error())
		}
		return utility.InternalServerError(c, "Failed to delete product")
	}
	return c.NoContent(http.StatusNoContent)
}
