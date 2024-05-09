package config

import (
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
)

type Config struct {
	Postgres_Port string `env:"POSTGRES_PORT"`
	Postgres_Host string `env:"POSTGRES_HOST"`
	Database_Name string `env:"POSTGRES_DB"`
	Postgres_User string `env:"POSTGRES_USER"`
	Postgres_Password string `env:"POSTGRES_PASSWORD"`
	Env string `env:"ENV"`
	Server_Address string `env: "SERVER_ADDR"`
	Server_Port string `env: "SERVER_PORT"`
	Timeout           time.Duration `env:"TIMEOUT"`
	IdleTimeout       time.Duration `env:"IDLE_TIMEOUT"`
}


func EnvLoad() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while loadig .env file: %v", err)
	}
	env := os.Getenv("ENV")
	log := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ldate)
	log.Printf("ENV is %s", env)
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("Failed to read config path")
	}
	log.Printf("CONFIG PATH is %s", configPath)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config does not exists: %s", err)
	}	
	var cfg Config
	return cfg
}
