package repository

import (
	"context"
	"tablelink-be-test/internal/domain/entity"
	"tablelink-be-test/internal/infrastructure/database"
)

type RoleRightRepositoryImpl struct {
	db *database.PostgresDB
}

func NewRoleRightRepositoryImpl(db *database.PostgresDB) *RoleRightRepositoryImpl {
	return &RoleRightRepositoryImpl{db: db}
}

func (r *RoleRightRepositoryImpl) GetByRoleIDAndRoute(ctx context.Context, roleID, route string) (*entity.RoleRight, error) {
	query := `SELECT id, role_id, section, route, r_create, r_read, r_update, r_delete FROM role_rights WHERE role_id = $1 AND route = $2`
	var roleRight entity.RoleRight
	row := r.db.DB.QueryRowContext(ctx, query, roleID, route)
	err := row.Scan(&roleRight.ID, &roleRight.RoleID, &roleRight.Section, &roleRight.Route, &roleRight.RCreate, &roleRight.RRead, &roleRight.RUpdate, &roleRight.RDelete)
	if err != nil {
		return nil, err
	}
	return &roleRight, nil
}
