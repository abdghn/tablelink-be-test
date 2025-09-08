package repository

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"tablelink-be-test/internal/domain/entity"
	"tablelink-be-test/internal/infrastructure/database"
	"time"
)

type UserRepositoryImpl struct {
	db *database.PostgresDB
}

func NewUserRepositoryImpl(db *database.PostgresDB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) GetByEmailAndPassword(ctx context.Context, email, password string) (*entity.User, error) {

	query := `SELECT id, role_id, name, email, password, created_at, updated_at, last_access FROM users WHERE email = $1`

	var user entity.User
	row := r.db.DB.QueryRowContext(ctx, query, email)

	err := row.Scan(&user.ID, &user.RoleID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.LastAccess)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) GetAll(ctx context.Context) ([]*entity.User, error) {
	query := `SELECT id, role_id, name, email, password, created_at, updated_at, last_access FROM users`

	rows, err := r.db.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.RoleID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.LastAccess)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *UserRepositoryImpl) GetByID(ctx context.Context, id string) (*entity.User, error) {
	query := `SELECT id, role_id, name, email, password, created_at, updated_at, last_access FROM users WHERE id = $1`
	var user entity.User

	row := r.db.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(&user.ID, &user.RoleID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.LastAccess)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) Create(ctx context.Context, user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()

	query := `INSERT INTO users (role_id, name, email, password, created_at, updated_at, last_access) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = r.db.DB.ExecContext(ctx, query, user.RoleID, user.Name, user.Email, string(hashedPassword), now, now, now)
	return err
}

func (r *UserRepositoryImpl) Update(ctx context.Context, user *entity.User) error {

	query := `UPDATE users SET name = $1, updated_at = $2 WHERE id = $3`
	_, err := r.db.DB.ExecContext(ctx, query, user.Name, time.Now(), user.ID)
	return err
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.DB.ExecContext(ctx, query, id)
	return err
}

func (r *UserRepositoryImpl) UpdateLastAccess(ctx context.Context, id string) error {
	query := `UPDATE users SET last_access = $1 WHERE id = $2`

	_, err := r.db.DB.ExecContext(ctx, query, time.Now(), id)
	return err
}
