package dto

import "time"

type CreateProductRequest struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"required,gt=0"`
	Stock int     `json:"stock" validate:"required,gte=0"`
}

type UpdateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price" validate:"gt=0"`
	Stock int     `json:"stock" validate:"gte=0"`
}

type ProductResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
