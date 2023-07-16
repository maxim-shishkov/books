package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/maxim-shishkov/books/pkg/config"
)

var DB *sql.DB

func ConnectDB() error {
	var err error
	DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Postgres.Host, config.Postgres.Port, config.Postgres.User, config.Postgres.Password, config.Postgres.DbName))
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	return nil
}

func Close() {
	DB.Close()
}
