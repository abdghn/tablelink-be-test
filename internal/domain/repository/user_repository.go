package repository

import (
	"context"
	"tablelink-be-test/internal/domain/entity"
)

type UserRepository interface {
	GetByEmailAndPassword(ctx context.Context, email, password string) (*entity.User, error)
	GetByID(ctx context.Context, id string) (*entity.User, error)
	GetAll(ctx context.Context) ([]*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id string) error
	UpdateLastAccess(ctx context.Context, id string) error
}

type RoleRepository interface {
	GetByID(ctx context.Context, id string) (*entity.Role, error)
}

type RoleRightRepository interface {
	GetByRoleIDAndRoute(ctx context.Context, roleID, route string) (*entity.RoleRight, error)
}
