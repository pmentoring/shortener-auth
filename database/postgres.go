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
)

func GetConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable",
		host, 5432, user, password, dbname)

	fmt.Println(psqlInfo)

	return sql.Open("postgres", psqlInfo)
}
