package main

import (
	"log"
	"os"

	"p3-graded-challenge-1-embapge/gateway-service/handler"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	e := echo.New()
	h := handler.NewGatewayHandler()

	// Proxy all /products* and /transactions* to shopping-service
	e.Any("/products*", h.ProxyToShoppingService)
	e.Any("/transactions*", h.ProxyToShoppingService)

	// Proxy all /payments* to payment-service
	e.Any("/payments*", h.ProxyToPaymentService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Println("Gateway running on port", port)
	e.Logger.Fatal(e.Start(":" + port))
}
