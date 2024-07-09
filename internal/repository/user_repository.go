package repository

import (
	"context"
	"database/sql"
	"os"

	"github.com/VadimRight/User_Microserver/internal/domain"
	"github.com/VadimRight/User_Microserver/internal/domain/entity"
)

type userRepository struct {
	Db *sql.DB
}

// Newpostgres.userRepository возвращает объект PostgresStorage
func NewUserRepository(db *sql.DB) domain.Repository {
	return &userRepository{Db: db}
}

// GetUserByUsername возвращает пользователя по его имени
func (s *userRepository) GetUserByUsername(ctx context.Context, username string) (entity.User, error) {
	var user entity.User
	data, err := os.ReadFile("../internal/repository/repositories_query/select_by_name.sql")
	query := string(data)
	err = s.Db.QueryRowContext(ctx, query, username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// UserCreate создает нового пользователя
func (s *userRepository) InsertUser(ctx context.Context, id entity.Uuid, username string, password string) (entity.User, error) {
	data, err := os.ReadFile("../internal/repository/repositories_query/insert.sql")
	query := string(data)
	_, err = s.Db.ExecContext(ctx, query, id, username, password)
	if err != nil {
		return entity.User{}, err
	}
	return entity.User{Id: id, Username: username}, nil
}

// GetUserByID возвращает пользователя по его ID
func (s *userRepository) GetUserByID(ctx context.Context, userID string) (entity.User, error) {
	var user entity.User
	data, err := os.ReadFile("../internal/repository/repositories_query/select.sql")
	query := string(data)
	err = s.Db.QueryRowContext(ctx, query, userID).Scan(&user.Id, &user.Username)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// GetAllUsers возвращает всех пользователей
func (s *userRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	data, err := os.ReadFile("../internal/repository/repositories_query/select.sql")
	query := string(data)
	rows, err := s.Db.QueryContext(ctx, query)
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
