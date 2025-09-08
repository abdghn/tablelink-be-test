package service

import (
	"context"
	"tablelink-be-test/internal/domain/entity"
)

type AuthService interface {
	Login(ctx context.Context, email, password string) (string, *entity.UserSession, error)
	Logout(ctx context.Context, token string) error
	ValidateToken(ctx context.Context, token string) (*entity.UserSession, error)
	ValidatePermission(ctx context.Context, userSession *entity.UserSession, route, method string) error
}

type CacheService interface {
	SetUserSession(ctx context.Context, token string, userSession *entity.UserSession) error
	GetUserSession(ctx context.Context, token string) (*entity.UserSession, error)
	DeleteUserSession(ctx context.Context, token string) error
}
