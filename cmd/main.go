package main


import (
	"log"
	"github.com/VadimRight/User_Microserver/internal/config"
	"github.com/VadimRight/User_Microserver/internal/postgres"
	"github.com/VadimRight/User_Microserver/api"
)

func main() {
	cfg := config.LoadConfig()
	database := postgres.InitPostgresDatabase()
	log.Printf("Database is initialized in %v", database)	
	api.InitServer(cfg)
}
