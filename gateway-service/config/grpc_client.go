package config

import (
	"context"
	"log"
	"net/url"
	"os"
	paymentpb "p3-graded-challenge-2-embapge/proto/payment"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClients struct {
	PaymentClient paymentpb.PaymentServiceClient
}

func NewGRPCClients() *GRPCClients {
	address := os.Getenv("PAYMENT_SERVICE_URL")
	if address == "" {
		address = "localhost:50052"
	}

	if u, err := url.Parse(address); err == nil && u.Host != "" {
		address = u.Host
	}

	paymentConn, err := grpc.DialContext(context.Background(), address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Payment Service: %v", err)
	}
	return &GRPCClients{
		PaymentClient: paymentpb.NewPaymentServiceClient(paymentConn),
	}
}
