package main

import (
	"log"
	"net"
	"os"
	"p3-graded-challenge-2-embapge/payment-service/config"
	"p3-graded-challenge-2-embapge/payment-service/internal/app"
	grpc2 "p3-graded-challenge-2-embapge/payment-service/internal/delivery/grpc"
	"p3-graded-challenge-2-embapge/payment-service/internal/infra"
	pb "p3-graded-challenge-2-embapge/proto/payment"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
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
	// paymentHandler := http.NewPaymentHandler(paymentApp)

	// e := echo.New()
	// e.POST("/payments", paymentHandler.CreatePayment)

	paymentGRPCHandler := grpc2.NewPaymentHandler(paymentApp)

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		log.Fatal("gRPC port is empty")
	}

	go func() {
		lis, err := net.Listen("tcp", ":"+grpcPort)
		if err != nil {
			log.Fatalf("Failed to listen gRPC: %v", err)
		}

		s := grpc.NewServer()

		pb.RegisterPaymentServiceServer(s, paymentGRPCHandler)
		log.Println("gRPC Auth Service running at:" + grpcPort)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Block main goroutine - ga langsung exit
	select {}
}
