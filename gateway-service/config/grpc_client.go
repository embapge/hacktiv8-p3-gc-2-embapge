package config

import (
	"log"
	"os"
	paymentpb "p3-graded-challenge-2-embapge/proto/payment"

	"google.golang.org/grpc"
)

type GRPCClients struct {
	PaymentClient paymentpb.PaymentServiceClient
}

func NewGRPCClients() *GRPCClients {
	address := os.Getenv("PAYMENT_SERVICE_URL")
	if address == "" {
		address = "localhost:50052"
	}
	paymentConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Auth Service: %v", err)
	}
	return &GRPCClients{
		PaymentClient: paymentpb.NewPaymentServiceClient(paymentConn),
	}
}
