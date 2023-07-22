package sqldb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env files")
	}

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PWD")
	sslMode := os.Getenv("PG_SSL_MODE")
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", dbUser, dbName, dbPassword, sslMode)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	DB = db
	return nil

}
