package main

import (
	"fmt"
	"log"
	"os"

	"auth-service/internal/auth/app"
	"auth-service/internal/auth/config"
	"auth-service/internal/auth/delivery/http"
	"auth-service/internal/auth/infra"
	"auth-service/pkg/hasher" // Shared logic (pkg)
	"auth-service/pkg/jwt"    // Shared logic (pkg)

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

	/*
		Matikan penggunaan MongoDB

		// 2. Init MongoDB connection with Connection Pooling (BEST PRACTICE)
		mongoURI := os.Getenv("MONGO_URI")
		dbName := os.Getenv("MONGO_DB")
		if mongoURI == "" || dbName == "" {
			log.Fatal("MONGO_URI & MONGO_DB must be set in .env")
		}

		// Baca pool size dari ENV (jika tidak ada pakai default)
		maxPool, _ := strconv.ParseUint(getEnvOrDefault("MONGO_POOL_MAX", "50"), 10, 64)
		minPool, _ := strconv.ParseUint(getEnvOrDefault("MONGO_POOL_MIN", "5"), 10, 64)

		clientOpts := options.Client().
			ApplyURI(mongoURI).
			SetMaxPoolSize(maxPool).
			SetMinPoolSize(minPool)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, clientOpts)
		if err != nil {
			log.Fatal("Gagal connect Mongo:", err)
		}
		defer func() {
			if err := client.Disconnect(ctx); err != nil {
				log.Println("Error disconnecting Mongo:", err)
			}
		}()
		log.Printf("✅ MongoDB Pooling aktif: min=%d max=%d\n", minPool, maxPool)
		userColl := client.Database(dbName).Collection("users")
		// 3. Init Infra Layer
		userRepo := &infra.MongoUserRepository{Collection: userColl}
	*/
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
	// }()

	// // 11. gRPC
	// grpcHandler := grpc2.NewAuthHandler(authApp)
	// authIntercptor := grpc2.NewAuthInterceptor(jwtManager)

	// // 12. Setup gRPC Server
	// grpcPort := os.Getenv("GRPC_PORT")
	// if grpcPort == "" {
	// 	log.Fatal("gRPC port is empty")
	// }

	// go func() {
	// 	lis, err := net.Listen("tcp", ":"+grpcPort)
	// 	if err != nil {
	// 		log.Fatalf("Failed to listen gRPC: %v", err)
	// 	}

	// 	s := grpc.NewServer(
	// 		grpc.UnaryInterceptor(authIntercptor.Unary()),
	// 	)

	// 	pb.RegisterAuthServiceServer(s, grpcHandler)
	// 	log.Println("gRPC Auth Service running at:" + grpcPort)
	// 	if err := s.Serve(lis); err != nil {
	// 		log.Fatalf("Failed to serve gRPC: %v", err)
	// 	}
	// }()

	// Block main goroutine - ga langsung exit
	// select {}
}

// Helper ambil ENV dengan fallback (biar clean)
func getEnvOrDefault(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
