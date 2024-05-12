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

var postgresCfg = config.LoadPostgresConfig()
var postgresUrl = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",postgresCfg.Postgres_Host, postgresCfg.Postgres_Port, postgresCfg.Postgres_User, postgresCfg.Postgres_Password, postgresCfg.Database_Name)	


func InitPostgresDatabase() *PostgresStorage {
	const op = "postgres.InitPostgresDatabase"
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
	defer db.Close()
	return &PostgresStorage{db: db}
}

func SaveNewUser(username, email, password string) {
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
