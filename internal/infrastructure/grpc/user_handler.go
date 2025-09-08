package grpc

import (
	"context"
	"errors"
	"strings"
	"tablelink-be-test/internal/application/usecase"
	userspb "tablelink-be-test/proto/users"
	"time"
)

type UserHandler struct {
	userspb.UnimplementedUserServiceServer
	userUsecase *usecase.UserUsecaseImpl
	authUsecase *usecase.AuthUsecaseImpl
}

func NewUserHandler(userUsecase *usecase.UserUsecaseImpl, authUsecase *usecase.AuthUsecaseImpl) *UserHandler {
	return &UserHandler{userUsecase: userUsecase, authUsecase: authUsecase}
}

func (h *UserHandler) validateRequest(ctx context.Context, token, serviceHeader, route, method string) error {

	if serviceHeader != "be" {
		return errors.New("invalid service header")
	}

	session, err := h.authUsecase.ValidateToken(ctx, token)
	if err != nil {
		return errors.New("invalid token")
	}

	return h.authUsecase.ValidatePermission(ctx, session, route, method)
}

func (h *UserHandler) GetAllUsers(ctx context.Context, req *userspb.GetAllUsersRequest) (*userspb.GetAllUserResponse, error) {
	token := strings.TrimPrefix(req.Token, "Bearer ")

	err := h.validateRequest(ctx, token, req.ServiceHeader, "/users/user", "GET")
	if err != nil {
		return &userspb.GetAllUserResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	session, _ := h.authUsecase.ValidateToken(ctx, token)

	return &userspb.GetAllUserResponse{
		Status:  true,
		Message: "Successfully",
		Data: &userspb.GetAllUsersData{
			User: &userspb.UserInfo{
				RoleId:     session.RoleID,
				RoleName:   session.RoleName,
				Name:       session.Name,
				Email:      session.Email,
				LastAccess: session.LastAccess.Format(time.RFC3339),
			},
		},
	}, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, req *userspb.CreateUserRequest) (*userspb.CreateUserResponse, error) {
	token := strings.TrimPrefix(req.Token, "Bearer ")

	err := h.validateRequest(ctx, token, req.ServiceHeader, "/users/user", "POST")
	if err != nil {
		return &userspb.CreateUserResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	err = h.userUsecase.CreateUser(ctx, req.RoleId, req.Name, req.Email, req.Password)
	if err != nil {
		return &userspb.CreateUserResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	return &userspb.CreateUserResponse{
		Status:  true,
		Message: "Successfully",
	}, nil

}

func (h *UserHandler) UpdateUser(ctx context.Context, req *userspb.UpdateUserRequest) (*userspb.UpdateUserResponse, error) {
	token := strings.TrimPrefix(req.Token, "Bearer ")

	err := h.validateRequest(ctx, token, req.ServiceHeader, "/users/user", "PUT")
	if err != nil {
		return &userspb.UpdateUserResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	err = h.userUsecase.UpdateUser(ctx, req.UserId, req.Name)
	if err != nil {
		return &userspb.UpdateUserResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	return &userspb.UpdateUserResponse{
		Status:  true,
		Message: "Successfully",
	}, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *userspb.DeleteUserRequest) (*userspb.DeleteUserResponse, error) {
	token := strings.TrimPrefix(req.Token, "Bearer ")
	err := h.validateRequest(ctx, token, req.ServiceHeader, "/users/user", "DELETE")
	if err != nil {
		return &userspb.DeleteUserResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	err = h.userUsecase.DeleteUser(ctx, req.UserId)
	if err != nil {
		return &userspb.DeleteUserResponse{
			Status:  false,
			Message: err.Error(),
		}, nil
	}

	return &userspb.DeleteUserResponse{
		Status:  true,
		Message: "Successfully",
	}, nil
}
