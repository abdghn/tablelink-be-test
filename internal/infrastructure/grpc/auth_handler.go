package grpc

import (
	"context"
	"tablelink-be-test/internal/application/usecase"
	authpb "tablelink-be-test/proto/auth"
)

type AuthHandler struct {
	authpb.UnimplementedAuthServiceServer
	authUsecase *usecase.AuthUsecaseImpl
}

func NewAuthHandler(authUsecase *usecase.AuthUsecaseImpl) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (a *AuthHandler) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	token, _, err := a.authUsecase.Login(ctx, req.Email, req.Password)
	if err != nil {
		return &authpb.LoginResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	return &authpb.LoginResponse{
		Status:  true,
		Message: "Successfully",
		Data: &authpb.LoginData{
			AccessToken: token,
		},
	}, nil
}

func (a *AuthHandler) Logout(ctx context.Context, req *authpb.LogoutRequest) (*authpb.LogoutResponse, error) {
	err := a.authUsecase.Logout(ctx, req.Token)
	if err != nil {
		return &authpb.LogoutResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	return &authpb.LogoutResponse{
		Status:  true,
		Message: "Successfully",
	}, nil
}
