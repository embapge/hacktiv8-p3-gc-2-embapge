package main

import (
	"log"
	"os"
	"p3-graded-challenge-1-embapge/payment-service/app"
	"p3-graded-challenge-1-embapge/payment-service/config"
	"p3-graded-challenge-1-embapge/payment-service/internal/delivery/http"
	"p3-graded-challenge-1-embapge/payment-service/internal/infra"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	db := config.ConnectDB()
	paymentCollection := config.GetCollection(db, "payments")

	paymentRepo := infra.NewPaymentRepository(paymentCollection)
	paymentApp := app.NewPaymentApp(paymentRepo)
	paymentHandler := http.NewPaymentHandler(paymentApp)

	e := echo.New()
	e.POST("/payments", paymentHandler.CreatePayment)

	e.Logger.Fatal(e.Start(":" + port))
}
