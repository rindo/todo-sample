package db

import (
	"fmt"
	"errors"
	"time"
	"database/sql"
	_ "github.com/lib/pq"
)

type Todo struct {
	Id	   string `json:id`
	Name	 string `json:name`
	Done	 bool `json:done`
	CreatedAt time.Time `created_at`
	UpdatedAt time.Time `updated_at`
}

func open() *sql.DB {
	db, err := sql.Open("postgres", "host=db port=5432 user=todo password=password dbname=todo sslmode=disable")

	if err != nil {
		fmt.Println("Error on opening database: %s", err)
	}

	return db
}

func Todos() ([]Todo, error) {
	db := open()

	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		fmt.Printf("Error on Todos(): %s", err)
		return []Todo{}, errors.New("DBエラー")
	}

	var todos []Todo
	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.Id, &todo.Name, &todo.Done, &todo.CreatedAt, &todo.UpdatedAt)
		todos = append(todos, todo)
	}

	defer db.Close()

	return todos, nil
}

func CreateTodo(name string) error {
	db := open()

	_, err := db.Exec("INSERT INTO todos (name) VALUES ($1)", name)
	if err != nil {
		fmt.Printf("Error on Todos(): %s", err)
		return errors.New("DBエラー")
	}

	return nil
}

func UpdateTodo(id int, name string, done bool) error {
	db := open()

	_, err := db.Exec("UPDATE todos SET name = $1, done = $2 WHERE id = $3", name, done, id)
	if err != nil {
		fmt.Printf("Error on Todos(): %s", err)
		return errors.New("DBエラー")
	}

	defer db.Close()

	return nil
}

func DeleteTodo(id int) error {
	db := open()

	_, err := db.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		fmt.Printf("Error on Todos(): %s", err)
		return errors.New("DBエラー")
	}

	defer db.Close()

	return nil
}
