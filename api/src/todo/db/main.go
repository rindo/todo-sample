package db

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

var instance *sql.DB

func GetInstance() *sql.DB {
	if instance == nil {
		instance, err := sql.Open("postgres", "host=db port=5432 user=todo password=password dbname=todo sslmode=disable")
		if err != nil {
			panic(err)
		}
		instance.SetMaxIdleConns(5)
		instance.SetMaxOpenConns(5)
	}

	return instance
}

func GetContext() context.Context {
	return context.Background()
}
