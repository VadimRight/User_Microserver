package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/VadimRight/User_Microserver/domain/entity"
	"github.com/VadimRight/User_Microserver/internal/config"
	"github.com/google/uuid"
)

type UserRepository struct {
	Db *sql.DB
}

// Newpostgres.UserRepository возвращает объект PostgresStorage
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func InitPostgresDatabase(cfg config.Config) UserRepository {
	const op = "postgres.InitPostgresDatabase"
	var postgresUrl = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Postgres.PostgresHost, cfg.Postgres.PostgresPort, cfg.Postgres.PostgresUser, cfg.Postgres.PostgresPassword, cfg.Postgres.DatabaseName)
	db, err := sql.Open("postgres", postgresUrl)
	if err != nil {
		log.Fatalf("Error while connecting to postgres database: %v", err)
	}

	createDatabase, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "user" (
		id UUID PRIMARY KEY,
		username VARCHAR(20) NOT NULL UNIQUE,
		email VARCHAR(20) NOT NULL UNIQUE,
		password CHAR(60) NOT NULL UNIQUE,
		is_verified BOOL NOT NULL DEFAULT false,
		is_activate BOOL NOT NULL DEFAULT false
	);`)
	if err != nil {
		log.Fatalf("%s: %v", op, err)
	}
	_, err = createDatabase.Exec()
	if err != nil {
		log.Fatalf("%s: %v", op, err)
	}
	return UserRepository{Db: db}
}

// ClosePostgres закрывает соединение с базой данных
func (s *UserRepository) ClosePostgres() error {
	return s.Db.Close()
}

// GetUserByUsername возвращает пользователя по его имени
func (s *UserRepository) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	err := s.Db.QueryRowContext(ctx, "SELECT id, username, password FROM users WHERE username=$1", username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UserCreate создает нового пользователя
func (s UserRepository) UserCreate(ctx context.Context, username string, password string) (*entity.User, error) {
	id := uuid.New()
	_, err := s.Db.ExecContext(ctx, "INSERT INTO users (id, username, password) VALUES ($1, $2, $3)", id, username, password)
	if err != nil {
		return nil, err
	}
	return &entity.User{Id: id, Username: username, Password: password}, nil
}

// GetUserByID возвращает пользователя по его ID
func (s UserRepository) GetUserByID(ctx context.Context, userID string) (*entity.User, error) {
	var user entity.User
	err := s.Db.QueryRowContext(ctx, "SELECT id, username FROM users WHERE id=$1", userID).Scan(&user.Id, &user.Username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllUsers возвращает всех пользователей
func (s UserRepository) GetAllUsers(ctx context.Context) ([]*entity.User, error) {
	rows, err := s.Db.QueryContext(ctx, "SELECT id, username FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.Id, &user.Username); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
