package repository_adapter

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type PgDB struct {
	db *sql.DB
}

func getEnvDB() string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("cannot locate env file")
	}

	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		panic("cannot load db hostname from env")
	}

	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		panic("cannot load db port from env")
	}

	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		panic("cannot load db user from env")
	}

	pass, ok := os.LookupEnv("DB_PASS")
	if !ok {
		panic("cannot load db user password from env")
	}

	dbname, ok := os.LookupEnv("DB_NAME")
	if !ok {
		panic("cannot load db name from env")
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname,
	)
}

func NewPgDB() *PgDB {
	db, err := sql.Open("postgres", getEnvDB())
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(1000)

	return &PgDB{db}
}
