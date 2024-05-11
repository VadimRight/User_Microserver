package postgres

import (
	"database/sql"
	"github.com/VadimRight/User_Microserver/internal/config"
	"log/slog"
)

func InitPostgresDatabase(cfg *config.Config) error {

	db, err := sql.Open("postgres", cfg)
}
