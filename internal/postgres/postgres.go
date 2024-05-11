package postgres

import (
	"database/sql"
	"github.com/VadimRight/User_Microserver/internal/config"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

func InitPostgresDatabase() *PostgresStorage {
	postgresCfg := config.LoadPostgresConfig()
	postgresUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",postgresCfg.Postgres_Host, postgresCfg.Postgres_Port, postgresCfg.Postgres_User, postgresCfg.Postgres_Password, postgresCfg.Database_Name)	
	db, err := sql.Open("postgres", postgresUrl)
	if err != nil {
		log.Fatalf("Error while connecting to postgres database: %v", err)	
	}

	createDatabase, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "user" (
		id UUID PRIMARY KEY,
		username VARCHAR(20) NOT NULL UNIQUE,
		email VARCHAR(20) NOT NULL UNIQUE,
		password CHAR(60) NOT NULL UNIQUE,
		is_verified BOOL NOT NULL DEFAULT false,
		is_activate BOOL NOT NULL DEFAULT false
	);`)
	if err != nil {
		log.Fatalf("Error while creating user table: %v", err)
	}
	_, err = createDatabase.Exec()
	if err != nil {
		log.Fatalf("Error while executing CREATE query: %v", err)
	}
	defer db.Close()
	return &PostgresStorage{db: db}
}
