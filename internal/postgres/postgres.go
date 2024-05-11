package postgres

import (
	"database/sql"
	"github.com/VadimRight/User_Microserver/internal/config"
	"fmt"
	"log"
)

type PostgresStorage struct {
	db *sql.DB
}

func InitPostgresDatabase() *PostgresStorage {
	postgresCfg := config.LoadPostgresConfig()
	postgresUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",postgresCfg.Postgres_Host, postgresCfg.Postgres_Port, postgresCfg.Postgres_User, postgresCfg.Postgres_Password, postgresCfg.Database_Name)
	db, err := sql.Open("postgres", postgresUrl)
	if err != nil {
		log.Fatalf("Error while connecting to postgres databse: %v", err)	
	}
	createDatabase, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS user(
		id UUID,
		username VARCHAR(20),
		email VARCHAR(20),
		password BINARY,
		is_verified BOOL DEFAULT 0
		is_activate BOOL DEFAULT 0
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
