package main

import (
	"log"
	"os"
	"p3-graded-challenge-1-embapge/payment-service/config"
	"p3-graded-challenge-1-embapge/payment-service/internal/handler"
	"p3-graded-challenge-1-embapge/payment-service/internal/repository"
	"p3-graded-challenge-1-embapge/payment-service/internal/router"
	"p3-graded-challenge-1-embapge/payment-service/internal/service"

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

	paymentRepo := repository.NewPaymentRepository(paymentCollection)
	paymentService := service.NewPaymentService(paymentRepo)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	e := echo.New()

	router.NewRouter(e, paymentHandler)

	e.Logger.Fatal(e.Start(":" + port))
}
