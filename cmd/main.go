package main


import (
	"log"
	"log/slog"
	"github.com/VadimRight/User_Microserver/pkg/prettylogger/handler/logger"
	"github.com/VadimRight/User_Microserver/internal/config"
	"github.com/VadimRight/User_Microserver/internal/postgres"
)

func main() {
	cfg := config.LoadConfig()
	logPretty := logger.SetupPrettyLogger(cfg.Env.Env)
	database := postgres.InitPostgresDatabase()
	logPretty.Info(
		"starting url-shortener",
		slog.String("env", cfg.Env.Env),
	)
	logPretty.Info(
		"this is log",
		slog.String("Postgres Name", cfg.Postgres.Database_Name),
		slog.String("Postgres Port", cfg.Postgres.Postgres_Port),
		slog.String("Postgres Host", cfg.Postgres.Postgres_Host),
		slog.String("Postgres User", cfg.Postgres.Postgres_User),
		slog.String("Server Port", cfg.Server.Server_Port),
		slog.String("Server_Address", cfg.Server.Server_Address),
		slog.Duration("Timeout", cfg.Server.Timeout),
		slog.Duration("Idle Timeout", cfg.Server.IdleTimeout),
	)
	log.Printf("Database is initialized in %v", database)	
}
