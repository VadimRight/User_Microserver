package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/VadimRight/User_Microserver/internal/config"
)

type PostgresStorage struct {
	Db *sql.DB
}

func InitPostgresDatabase(cfg config.Config) PostgresStorage {
	const op = "postgres.InitPostgresDatabase"
	var postgresUrl = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Postgres.PostgresHost, cfg.Postgres.PostgresPort, cfg.Postgres.PostgresUser, cfg.Postgres.PostgresPassword, cfg.Postgres.DatabaseName)
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
		log.Fatalf("%s: %v", op, err)
	}
	_, err = createDatabase.Exec()
	if err != nil {
		log.Fatalf("%s: %v", op, err)
	}
	return PostgresStorage{Db: db}
}

// ClosePostgres закрывает соединение с базой данных
func (s *PostgresStorage) ClosePostgres() error {
	return s.Db.Close()
}

// Newpostgres.PostgresStorage возвращает объект PostgresStorage
func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{Db: db}
}
