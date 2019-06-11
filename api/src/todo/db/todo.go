package db

import (
	"errors"
	. "todo/models"
	"github.com/volatiletech/sqlboiler/boil"
)

func GetTodos() (TodoSlice, error) {
	todos, err := Todos().All(GetContext(), GetInstance())
	if err != nil {
		return todos, err
	}
	
	return todos, nil
}

func CreateTodo(name string) error {
	var todo Todo
	todo.Name = name

	err := todo.Insert(GetContext(), GetInstance(), boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func UpdateTodo(id int, name string, done bool) error {
	todo, err := FindTodo(GetContext(), GetInstance(), id)

	if todo == nil {
		return errors.New("not found")
	}
	if err != nil {
		return err
	}

	todo.Name = name
	todo.Done = done
	_, err = todo.Update(GetContext(), GetInstance(), boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func DeleteTodo(id int) error {
	todo, err := FindTodo(GetContext(), GetInstance(), id)

	if todo == nil {
		return errors.New("not found")
	}
	if err != nil {
		return err
	}

	_, err = todo.Delete(GetContext(), GetInstance())
	if err != nil {
		return err
	}

	return nil
}
