package main

import (
	"log"
	"os"
	"p3-graded-challenge-1-embapge/shopping-service/config"
	"p3-graded-challenge-1-embapge/shopping-service/internal/handler"
	"p3-graded-challenge-1-embapge/shopping-service/internal/repository"
	"p3-graded-challenge-1-embapge/shopping-service/internal/router"
	"p3-graded-challenge-1-embapge/shopping-service/internal/scheduler"
	"p3-graded-challenge-1-embapge/shopping-service/internal/service"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// @title Shopping Service API
// @version 1.0
// @description This is a sample server for a shopping service.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
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
	productCollection := config.GetCollection(db, "products")
	transactionCollection := config.GetCollection(db, "transactions")

	productRepo := repository.NewProductRepository(productCollection)
	transactionRepo := repository.NewTransactionRepository(transactionCollection)

	productService := service.NewProductService(productRepo)
	transactionService := service.NewTransactionService(transactionRepo, productRepo)

	productHandler := handler.NewProductHandler(productService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	// Start scheduler for auto-fail payment > 3 jam
	scheduler.StartPaymentFailScheduler(transactionRepo)

	e := echo.New()

	router.NewRouter(e, productHandler, transactionHandler)

	e.Logger.Fatal(e.Start(":" + port))
}
