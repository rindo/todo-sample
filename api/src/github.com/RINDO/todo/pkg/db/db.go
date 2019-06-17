package db

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

var instance *sql.DB

func GetInstance() *sql.DB {
	if instance == nil {
		conn, err := sql.Open("postgres", "host=db port=5432 user=todo password=password dbname=todo sslmode=disable")
		if err != nil {
			panic(err)
		}
		conn.SetMaxIdleConns(5)
		conn.SetMaxOpenConns(5)
		instance = conn
	}

	return instance
}

func GetContext() context.Context {
	return context.Background()
}
