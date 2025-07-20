package grpc

import (
	"context"
	"p3-graded-challenge-2-embapge/payment-service/internal/app"
	"p3-graded-challenge-2-embapge/payment-service/internal/delivery/http/dto"
	pb "p3-graded-challenge-2-embapge/proto/payment"

	// If the above import path is incorrect, update it to the correct relative or module path, for example:
	// pb "p3-graded-challenge-2-embapge/payment-service/internal/proto/payment"
	// or
	// pb "payment-service/proto/payment"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PaymentHander struct {
	pb.UnimplementedPaymentServiceServer
	app *app.PaymentApp
}

func NewPaymentHandler(app *app.PaymentApp) *PaymentHander {
	return &PaymentHander{app: app}
}

func (p *PaymentHander) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.PaymentResponse, error) {
	createPaymentRequest := dto.CreatePaymentRequest{
		Amount: float64(req.Amount),
	}

	payment, err := p.app.CreatePayment(ctx, createPaymentRequest)
	if err != nil {
		return nil, errToGRPC(err, 13)
	}

	return &pb.PaymentResponse{
		Id:        payment.ID,
		Amount:    float32(payment.Amount),
		Status:    payment.Status,
		CreatedAt: timestamppb.New(payment.CreatedAt),
		UpdatedAt: timestamppb.New(payment.UpdatedAt),
	}, nil
}

func (p *PaymentHander) GetAllPayment(ctx context.Context, _ *emptypb.Empty) (*pb.ListPaymentResponse, error) {
	payments, err := p.app.GetAllPayment(ctx)
	if err != nil {
		return nil, errToGRPC(err, 13)
	}

	var grpcPayments []*pb.PaymentResponse
	for _, payment := range *payments {
		grpcPayment := pb.PaymentResponse{
			Id:        payment.ID.Hex(),
			Amount:    float32(payment.Amount),
			Status:    payment.Status,
			CreatedAt: timestamppb.New(payment.CreatedAt),
			UpdatedAt: timestamppb.New(payment.UpdatedAt),
		}
		grpcPayments = append(grpcPayments, &grpcPayment)
	}

	return &pb.ListPaymentResponse{Payments: grpcPayments}, nil
}

func (p *PaymentHander) GetByIDPayment(ctx context.Context, req *pb.PaymentIDRequest) (*pb.PaymentResponse, error) {
	payment, err := p.app.FindByIDPayment(ctx, req.GetId())
	if err != nil {
		return nil, errToGRPC(err, 13)
	}

	grpcPayment := pb.PaymentResponse{
		Id:        payment.ID,
		Amount:    float32(payment.Amount),
		Status:    payment.Status,
		CreatedAt: timestamppb.New(payment.CreatedAt),
		UpdatedAt: timestamppb.New(payment.UpdatedAt),
	}

	return &grpcPayment, nil
}

func (p *PaymentHander) DeletePayment(ctx context.Context, req *pb.PaymentIDRequest) (*pb.PaymentMessageResponse, error) {
	err := p.app.DeletePayment(ctx, req.GetId())
	if err != nil {
		return nil, errToGRPC(err, 13)
	}

	return &pb.PaymentMessageResponse{
		Message: "Payment deleted successfully",
	}, nil
}
