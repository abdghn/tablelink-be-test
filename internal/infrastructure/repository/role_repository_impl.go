package repository

import (
	"context"
	"tablelink-be-test/internal/domain/entity"
	"tablelink-be-test/internal/infrastructure/database"
)

type RoleRepositoryImpl struct {
	db *database.PostgresDB
}

func NewRoleRepositoryImpl(db *database.PostgresDB) *RoleRepositoryImpl {
	return &RoleRepositoryImpl{db: db}
}

func (r *RoleRepositoryImpl) GetByID(ctx context.Context, id string) (*entity.Role, error) {
	query := `SELECT id, name FROM roles WHERE id = $1`
	var role entity.Role
	row := r.db.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&role.ID, &role.Name)
	if err != nil {
		return nil, err
	}
	return &role, nil
}
