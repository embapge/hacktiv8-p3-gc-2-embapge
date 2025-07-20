package grpc

import (
	"context"
	"p3-graded-challenge-2-embapge/auth-service/internal/auth/app"
	"p3-graded-challenge-2-embapge/auth-service/internal/auth/delivery/grpc/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// handler grpc

type AuthGRPCHandler struct {
	pb.UnimplementedAuthServiceServer
	App *app.AuthApp
}

// Construct
func NewAuthHandler(app *app.AuthApp) *AuthGRPCHandler {
	return &AuthGRPCHandler{App: app}
}

// Register User via GRPC

func (h *AuthGRPCHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	user, err := h.App.Register(ctx, req.GetName(), req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, errToGRPC(err)
	}

	token, _ := h.App.JWTManager.GenerateToken(user.ID, user.Email)
	return &pb.AuthResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}, nil
}

// Login

func (h *AuthGRPCHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	token, err := h.App.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, errToGRPC(err)
	}

	return &pb.AuthResponse{
		Email: req.GetEmail(),
		Token: token,
	}, nil
}

// helper to GRPC
func errToGRPC(err error) error {
	switch err {
	case app.ErrValidation:
		return status.Error(codes.InvalidArgument, err.Error())
	case app.ErrEmailExist:
		return status.Error(codes.AlreadyExists, err.Error())
	case app.ErrUnauthorized:
		return status.Error(codes.Unauthenticated, err.Error())
	default:
		return status.Error(codes.Internal, err.Error())
	}
}
