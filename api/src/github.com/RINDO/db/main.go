package main

import (
	"fmt"
	"time"
	"database/sql"
	_ "github.com/lib/pq"
)

type Todo struct {
	Id 	      string
	Name      string
	Done 	  bool
	CreatedAt time.Time
	UpdatedAt tmie.Time
}

func open() {
	db, err := sql.Open("postgres", "host=db port=5432 user=todo password=password dbname=todo")

    if err != nil {
        fmt.Println(err)
        panic()
    }

    return db
}

func Todos() {
    db := open()

    defer db.Close()

    err := db.Query("SELECT * FROM todos")

    if err != nil {
        fmt.Printf(err)
    }

    var todos[]Todo
    for rows.Next() {
        var todo Todo
        rows.Scan(
            &todo.Id, 
            &todo.Name, 
            &todo.Done, 
            &todo.CreatedAt,
            &todo.UpdatedAt
        )
        todos = append(todos, todo)
    }

    return todos
}