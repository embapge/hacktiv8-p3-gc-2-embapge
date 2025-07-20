package main

import (
	"fmt"
	"log"
	"os"

	"auth-service/internal/auth/app"
	"auth-service/internal/auth/config"
	"auth-service/internal/auth/delivery/http"
	"auth-service/internal/auth/infra"
	"auth-service/pkg/hasher"
	"auth-service/pkg/jwt"

	_ "auth-service/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Auth Service API
// @version 1.0
// @description Microservices Auth (DDD + Mongo + JWT) for Tokokecil
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	// 1. Load env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db := config.MySQLInit()
	userRepo := infra.NewMySQLUserRepository(db)

	// 4. Init Shared Logic dari pkg
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET must be set in .env")
	}
	jwtManager := jwt.NewManager(jwtSecret)
	passwordHasher := hasher.NewBcrypt()

	// 5. Init Application Layer
	authApp := app.NewAuthApp(userRepo, passwordHasher, jwtManager)

	// go func() {
	// 6. Init Delivery Layer (HTTP handler)
	authHandler := http.NewAuthHandler(authApp)

	// 7. Setup Echo (Router)
	e := echo.New()

	// 8. Swagger docs route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// 9. Auth endpoints
	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)

	// 10. Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Auth Service running at http://localhost:" + port)
	fmt.Println("Swagger UI: http://localhost:" + port + "/swagger/index.html")
	if err := e.Start(":" + port); err != nil {
		log.Fatal(err)
	}
}
