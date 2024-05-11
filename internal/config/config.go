package config

import (
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
)

type Config struct {
}

type EnvConfig struct {
	Env string 
	EnvPath string
}

type PostgresConfig struct {	
	Postgres_Port string 
	Postgres_Host string 
	Database_Name string 
	Postgres_User string 
	Postgres_Password string 
}

type ServerConfig struct {
	Server_Address string 
	Server_Port string 
	Timeout           time.Duration 
	IdleTimeout       time.Duration 
}

func LoadEnvConfig() *EnvConfig {
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
	var cfg EnvConfig
	return &cfg
}

func LoadPostgresConfig() *PostgresConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading postgres env: %v", err)
	}
	postgresPort, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		log.Fatal("Can't read POSTGRES_PORT")
	}
	postgresHost, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		log.Fatal("Can't read POSTGRES_HOST")
	}
	postgresPassword, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		log.Fatal()
	}
	postgresDB, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		log.Fatal("Can't read POSTGRES_DB")
	}
	postgresUser, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		log.Fatal("Can't read POSTGRES_USER")
	}
	return &PostgresConfig {
		Postgres_Port: postgresPort,
		Postgres_Host: postgresHost,
		Database_Name: postgresDB,
		Postgres_User: postgresUser,
		Postgres_Password: postgresPassword,	
	}
}

func LoadServerConfig() *ServerConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err while loading server env: %v", err)
	}
	serverPort, ok := os.LookupEnv("SERVER_PORT")
	if !ok {
		log.Fatal("Can't read SERVER_PORT")
	}
	serverAddr, ok := os.LookupEnv("SERVER_ADDR")
	if !ok {
		log.Fatal("Can't read SERVER_ADDR")
	}
	timeout, ok := os.LookupEnv("TIMEOUT")
	if !ok {
		log.Fatal("Can't read TIMEOUT")
	}
	timeoutTime, err := time.ParseDuration(timeout)
	if err != nil {
		log.Fatalf("error while parsing timeout")
	}
	idleTimeout, ok := os.LookupEnv("IDLE_TIMEOUT")
	if !ok {
		log.Fatal("Can't read IDLE_TIMEOUT")
	}
	idleTimeoutTime, err := time.ParseDuration(idleTimeout)
	if err != nil {
		log.Fatalf("error while parsing idle time")
	}
	return &ServerConfig {
		Server_Address: serverAddr, 
		Server_Port: serverPort,
		Timeout: timeoutTime, 
		IdleTimeout: idleTimeoutTime,  	
	}
}
