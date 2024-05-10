package config

import (
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Postgres_Port string `env:"POSTGRES_PORT"`
	Postgres_Host string `env:"POSTGRES_HOST"`
	Database_Name string `env:"POSTGRES_DB"`
	Postgres_User string `env:"POSTGRES_USER"`
	Postgres_Password string `env:"POSTGRES_PASSWORD"`
	Env string `env:"ENV"`
	Server_Address string `env:"SERVER_ADDR"`
	Server_Port string `env:"SERVER_PORT"`
	Timeout           time.Duration `env:"TIMEOUT"`
	IdleTimeout       time.Duration `env:"IDLE_TIMEOUT"`
}

func EnvLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}		
	env := os.Getenv("ENV")
	log := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	log.Printf("ENV is %s", env)
	configPath := os.Getenv("CONFIG_PATH")
	log.Printf("CONFIG_PATH is %s", configPath)
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exists", configPath)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatal("cannot read database config")
	}
	return &cfg
}
