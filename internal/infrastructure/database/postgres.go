package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB(host, port, user, password, dbname string) (*PostgresDB, error) {
	psql := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDB{DB: db}, nil
}

func (p *PostgresDB) Close() error {
	return p.DB.Close()
}

const createTablesSQL = `
CREATE TABLE IF NOT EXISTS roles (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(255) PRIMARY KEY,
    role_id VARCHAR(255) REFERENCES roles(id),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_access TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS role_rights (
    id VARCHAR(255) PRIMARY KEY,
    role_id VARCHAR(255) REFERENCES roles(id),
    section VARCHAR(255) NOT NULL,
    route VARCHAR(255) NOT NULL,
    r_create INT DEFAULT 0,
    r_read INT DEFAULT 0,
    r_update INT DEFAULT 0,
    r_delete INT DEFAULT 0
);

-- Insert default data
INSERT INTO roles (id, name) VALUES ('1', 'Admin') ON CONFLICT (id) DO NOTHING;
INSERT INTO roles (id, name) VALUES ('2', 'User') ON CONFLICT (id) DO NOTHING;

INSERT INTO users (id, role_id, name, email, password) 
VALUES ('1', '1', 'Administrator', 'admin@gmail.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi') 
ON CONFLICT (id) DO NOTHING;

INSERT INTO role_rights (id, role_id, section, route, r_create, r_read, r_update, r_delete)
VALUES 
('1', '1', 'be', '/users/user', 1, 1, 1, 1),
('2', '2', 'be', '/users/user', 0, 1, 0, 0)
ON CONFLICT (id) DO NOTHING;
`

func (p *PostgresDB) CreateTables() error {
	_, err := p.DB.Exec(createTablesSQL)
	return err
}
