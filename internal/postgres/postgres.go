package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/VadimRight/User_Microserver/internal/config"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
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
	return PostgresStorage{db: db}
}

// ClosePostgres закрывает соединение с базой данных
func (s *PostgresStorage) ClosePostgres() error {
	return s.db.Close()
}

func RegisterUser(username, email, password string) {
	var postgresCfg = config.LoadPostgresConfig()
	var postgresUrl = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", postgresCfg.PostgresHost, postgresCfg.PostgresPort, postgresCfg.PostgresUser, postgresCfg.PostgresPassword, postgresCfg.DatabaseName)
	const op = "postgres.SaveNewUser"
	db, err := sql.Open("postgres", postgresUrl)
	if err != nil {
		log.Fatalf("%s: %v", op, err)
	}
	uuidUserId := uuid.New()
	createNewUser, err := db.Prepare(`
	INSERT INTO "user" (id, username, email, password, is_verified, is_activate)
	VALUE (?, ?, ?, ?);`)
	if err != nil {
		log.Fatalf("%s: %v", op, err)
	}
	res, err := createNewUser.Exec(uuidUserId, username, email, password)
	if err != nil {
		if postgresErr, ok := err.(*pq.Error); ok {
			fmt.Println("pq error:", postgresErr.Code.Name())
		}
	}
	_, err = res.LastInsertId()
	if err != nil {
		log.Fatalf("%s failed to get last inserted id: %v", op, err)
	}
}

func GetUser(username string) {
	var postgresCfg = config.LoadPostgresConfig()
	var postgresUrl = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", postgresCfg.PostgresHost, postgresCfg.PostgresPort, postgresCfg.PostgresUser, postgresCfg.PostgresPassword, postgresCfg.DatabaseName)
	const op = "postgres.GetUser"
	db, err := sql.Open("postgres", postgresUrl)
	if err != nil {
		log.Fatalf("%s: %v", op, err)
	}
	getUser, err := db.Prepare(`
		SELECT id, username, email, is_verified, is_activate FROM "user" WHERE id = ?;
	`)
	if err != nil {
		log.Fatalf("%s: %v", op, err)
	}
	_, err = getUser.Exec(username)
	if err != nil {
		if postgresErr, ok := err.(*pq.Error); ok {
			fmt.Println("pq error:", postgresErr.Code.Name())
		}
	}
}
