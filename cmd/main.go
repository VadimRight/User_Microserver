package main


import (
	"github.com/VadimRight/User_Microserver/internal/config"
	"github.com/VadimRight/User_Microserver/internal/postgres"
	"github.com/VadimRight/User_Microserver/api"
)

func main() {
	cfg := config.LoadConfig()
	_ = postgres.InitPostgresDatabase()
	api.InitServer(cfg)
}
