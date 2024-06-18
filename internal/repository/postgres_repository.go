package repository

import (
	"context"

	"github.com/VadimRight/User_Microserver/internal/postgres"
	"github.com/google/uuid"
)

// GetUserByUsername возвращает пользователя по его имени
func (s *postgres.PostgresStorage) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := s.DB.QueryRowContext(ctx, "SELECT id, username, password FROM users WHERE username=$1", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UserCreate создает нового пользователя
func (s *postgres.PostgresStorage) UserCreate(ctx context.Context, username string, password string) (*model.User, error) {
	id := uuid.New().String()
	_, err := s.DB.ExecContext(ctx, "INSERT INTO users (id, username, password) VALUES ($1, $2, $3)", id, username, password)
	if err != nil {
		return nil, err
	}
	return &model.User{ID: id, Username: username, Password: password}, nil
}

// GetUserByID возвращает пользователя по его ID
func (s *postgres.PostgresStorage) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	var user model.User
	err := s.DB.QueryRowContext(ctx, "SELECT id, username FROM users WHERE id=$1", userID).Scan(&user.ID, &user.Username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllUsers возвращает всех пользователей
func (s *postgres.PostgresStorage) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	rows, err := s.DB.QueryContext(ctx, "SELECT id, username FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
