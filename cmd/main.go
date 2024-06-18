package main

import (
	"github.com/VadimRight/User_Microserver/api"
	"github.com/VadimRight/User_Microserver/internal/config"
	"github.com/VadimRight/User_Microserver/internal/postgres"
)

func main() {
	cfg := config.LoadConfig()
	postgres.InitPostgresDatabase(cfg)
	api.InitServer(cfg)
}
