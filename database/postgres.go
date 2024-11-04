package database

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx"
	_ "github.com/lib/pq"
	"os"
)

var (
	host     = os.Getenv("POSTGRES_HOST")
	user     = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname   = os.Getenv("POSTGRES_DB")
	testDb   = os.Getenv("POSTGRES_TEST_DB")
)

func GetConnection() (*sql.DB, error) {
	activeDb := dbname
	if os.Getenv("GO_ENV") == "test" {
		activeDb = testDb
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, 5432, user, password, activeDb)
	fmt.Println(psqlInfo)

	return sql.Open("postgres", psqlInfo)
}
