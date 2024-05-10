package main

import (
	"log/slog"
	"github.com/VadimRight/User_Microserver/pkg/prettylogger/handler/logger"
	"github.com/VadimRight/User_Microserver/internal/config"
)

func main() {
	cfg := config.EnvLoad()
	log := slogpretty.SetupPrettyLogger(cfg.Env)
	log.Info(
		"starting url-shortener",
		slog.String("env", cfg.Env),
	)
	log.Info(
		"this is log",
		slog.String("Postgres Name", cfg.Database_Name),
		slog.String("Postgres Port", cfg.Postgres_Port),
		slog.String("Postgres Host", cfg.Postgres_Host),
		slog.String("Postgres User", cfg.Postgres_User),
		slog.String("Server Port", cfg.Server_Port),
		slog.Duration("Timeout", cfg.Timeout),
		slog.Duration("Idle Timeout", cfg.IdleTimeout),
	)
}
