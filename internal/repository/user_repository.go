package repository

import (
	"context"
	"database/sql"

	"github.com/VadimRight/User_Microserver/domain/entity"
)

type UserRepository struct {
	Db *sql.DB
}

// Newpostgres.UserRepository возвращает объект PostgresStorage
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

// GetUserByUsername возвращает пользователя по его имени
func (s UserRepository) GetUserByUsername(ctx context.Context, username string) (entity.User, error) {
	var user entity.User
	err := s.Db.QueryRowContext(ctx, "SELECT id, username, password FROM users WHERE username=$1", username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// UserCreate создает нового пользователя
func (s UserRepository) InsertUser(ctx context.Context, id entity.Uuid, username string, password string) (entity.User, error) {
	_, err := s.Db.ExecContext(ctx, "INSERT INTO users (id, username, password) VALUES ($1, $2, $3)", id, username, password)
	if err != nil {
		return entity.User{}, err
	}
	return entity.User{Id: id, Username: username}, nil
}

// GetUserByID возвращает пользователя по его ID
func (s UserRepository) GetUserByID(ctx context.Context, userID string) (entity.User, error) {
	var user entity.User
	err := s.Db.QueryRowContext(ctx, "SELECT id, username FROM users WHERE id=$1", userID).Scan(&user.Id, &user.Username)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// GetAllUsers возвращает всех пользователей
func (s UserRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	rows, err := s.Db.QueryContext(ctx, "SELECT id, username FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(user.Id, user.Username); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
