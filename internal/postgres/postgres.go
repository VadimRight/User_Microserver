package postgres

import (
	"database/sql"
	"github.com/VadimRight/User_Microserver/internal/config"
	"github.com/VadimRight/User_Microserver/pkg/prettylogger/handler/logger"
)

func InitPostgresDatabase(cfg *config.Config) error {
	log := logger.SetupPrettyLogger("host=%s port=%s user=%s password=%s dbname=%s")	
}
