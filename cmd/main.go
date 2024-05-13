package main


import (
	"log"
	"log/slog"
	"github.com/VadimRight/User_Microserver/pkg/prettylogger/handler/logger"
	"github.com/VadimRight/User_Microserver/internal/config"
	"github.com/VadimRight/User_Microserver/internal/postgres"
	"github.com/VadimRight/User_Microserver/api"
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
		slog.String("Postgres Name", cfg.Postgres.DatabaseName),
		slog.String("Postgres Port", cfg.Postgres.PostgresPort),
		slog.String("Postgres Host", cfg.Postgres.PostgresHost),
		slog.String("Postgres User", cfg.Postgres.PostgresUser),
		slog.String("Server Port", cfg.Server.ServerPort),
		slog.String("Server_Address", cfg.Server.ServerAddress),
		slog.Duration("Timeout", cfg.Server.Timeout),
		slog.Duration("Idle Timeout", cfg.Server.IdleTimeout),
	)
	log.Printf("Database is initialized in %v", database)	
	api.InitServer(cfg)
}
