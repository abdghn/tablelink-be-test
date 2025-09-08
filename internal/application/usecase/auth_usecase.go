package usecase

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"tablelink-be-test/internal/domain/entity"
	"tablelink-be-test/internal/domain/repository"
	"tablelink-be-test/internal/domain/service"
	"time"
)

type AuthUsecaseImpl struct {
	userRepo      repository.UserRepository
	roleRepo      repository.RoleRepository
	roleRightRepo repository.RoleRightRepository
	cacheService  service.CacheService
}

func NewAuthUsecaseImpl(userRepo repository.UserRepository, roleRepo repository.RoleRepository, roleRightRepo repository.RoleRightRepository, cacheService service.CacheService) *AuthUsecaseImpl {
	return &AuthUsecaseImpl{userRepo: userRepo, roleRepo: roleRepo, roleRightRepo: roleRightRepo, cacheService: cacheService}
}

func (a *AuthUsecaseImpl) Login(ctx context.Context, email, password string) (string, *entity.UserSession, error) {
	user, err := a.userRepo.GetByEmailAndPassword(ctx, email, password)
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	role, err := a.roleRepo.GetByID(ctx, user.RoleID)
	if err != nil {
		return "", nil, err
	}

	token, err := generateToken()
	if err != nil {
		return "", nil, err
	}

	err = a.userRepo.UpdateLastAccess(ctx, user.ID)
	if err != nil {
		return "", nil, err
	}

	userSession := &entity.UserSession{
		UserID:     user.ID,
		RoleID:     role.ID,
		RoleName:   role.Name,
		Name:       user.Name,
		Email:      user.Email,
		LastAccess: time.Now(),
	}

	err = a.cacheService.SetUserSession(ctx, token, userSession)
	if err != nil {
		return "", nil, err
	}

	return token, userSession, nil
}

func (a *AuthUsecaseImpl) Logout(ctx context.Context, token string) error {
	return a.cacheService.DeleteUserSession(ctx, token)
}

func (a *AuthUsecaseImpl) ValidateToken(ctx context.Context, token string) (*entity.UserSession, error) {
	return a.cacheService.GetUserSession(ctx, token)
}

func (a *AuthUsecaseImpl) ValidatePermission(ctx context.Context, userSession *entity.UserSession, route, method string) error {
	roleRight, err := a.roleRightRepo.GetByRoleIDAndRoute(ctx, userSession.RoleID, route)
	if err != nil {
		return errors.New("access denied")
	}

	if roleRight.Section != "be" {
		return errors.New("invalid section")
	}

	switch method {
	case "GET":
		if roleRight.RRead != 1 {
			return errors.New("read access denied")
		}
	case "POST":
		if roleRight.RCreate != 1 {
			return errors.New("create access denied")
		}
	case "PUT":
		if roleRight.RUpdate != 1 {
			return errors.New("update access denied")
		}
	case "DELETE":
		if roleRight.RDelete != 1 {
			return errors.New("delete access denied")
		}
	default:
		return errors.New("invalid method")
	}

	return nil
}

func generateToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
