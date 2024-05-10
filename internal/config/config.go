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



func EnvLoad() *Config {
  log := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ldate)

  err := godotenv.Load(".env")
  if err != nil {
    log.Fatalf("Error while loadig .env file: %v", err)
  }

  env, ok := os.LookupEnv("ENV")
  if !ok {
    log.Fatal("Failed to read ENV")
  }

  configPath, ok := os.LookupEnv("CONFIG_PATH")
  if !ok {
    log.Fatal("Failed to read CONFIG_PATH")
  }

  _, err = os.Stat(configPath)
  if os.IsNotExist(err) {
    log.Fatalf("config does not exists: %s", err)
  }

  log.Println(configPath)

  postgresPort, ok := os.LookupEnv("POSTGRES_PORT")
  if !ok {
    log.Fatal("Failed to read POSTGRES_PORT")
  }

  return &Config{
    Postgres_Port:     postgresPort,
    Postgres_Host:     "",
    Database_Name:     "",
    Postgres_User:     "",
    Postgres_Password: "",
    Env:               env,
    Server_Address:    "",
    Server_Port:       "",
    Timeout:           0,
    IdleTimeout:       0,
  }
}
