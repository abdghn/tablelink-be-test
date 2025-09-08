package usecase

import (
	"context"
	"github.com/google/uuid"
	"tablelink-be-test/internal/domain/entity"
	"tablelink-be-test/internal/domain/repository"
)

type UserUsecaseImpl struct {
	userRepo repository.UserRepository
	roleRepo repository.RoleRepository
}

func NewUserUsecaseImpl(userRepo repository.UserRepository, roleRepo repository.RoleRepository) *UserUsecaseImpl {
	return &UserUsecaseImpl{userRepo: userRepo, roleRepo: roleRepo}
}

func (u *UserUsecaseImpl) GetAll(ctx context.Context) ([]*entity.User, error) {
	return u.userRepo.GetAll(ctx)
}

func (u *UserUsecaseImpl) GetByID(ctx context.Context, id string) (*entity.User, error) {
	return u.userRepo.GetByID(ctx, id)
}

func (u *UserUsecaseImpl) CreateUser(ctx context.Context, roleID, name, email, password string) error {
	user := &entity.User{
		ID:       uuid.New().String(),
		RoleID:   roleID,
		Name:     name,
		Email:    email,
		Password: password,
	}

	return u.userRepo.Create(ctx, user)
}

func (u *UserUsecaseImpl) UpdateUser(ctx context.Context, id, name string) error {
	user := &entity.User{
		Name: name,
	}

	return u.userRepo.Update(ctx, user)
}

func (u *UserUsecaseImpl) DeleteUser(ctx context.Context, id string) error {
	return u.userRepo.Delete(ctx, id)
}
