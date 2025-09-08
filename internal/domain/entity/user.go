package entity

import "time"

type User struct {
	ID         string `db:"id" json:"id"`
	RoleID     string `db:"role_id" json:"role_id"`
	Name       string `db:"name" json:"name"`
	Email      string `db:"email" json:"email"`
	Password   string `db:"password" json:"password"`
	CreatedAt  string `db:"created_at" json:"created_at"`
	UpdatedAt  string `db:"updated_at" json:"updated_at"`
	LastAccess string `db:"last_access" json:"last_access"`
}

type Role struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type RoleRight struct {
	ID      string `db:"id" json:"id"`
	RoleID  string `db:"role_id" json:"role_id"`
	Section string `db:"section" json:"section"`
	Route   string `db:"route" json:"route"`
	RCreate int    `db:"r_create" json:"r_create"`
	RRead   int    `db:"r_read" json:"r_read"`
	RUpdate int    `db:"r_update" json:"r_update"`
	RDelete int    `db:"r_delete" json:"r_delete"`
}

type UserSession struct {
	UserID     string    `json:"user_id"`
	RoleID     string    `json:"role_id"`
	RoleName   string    `json:"role_name"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	LastAccess time.Time `json:"last_access"`
}
