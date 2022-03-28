package db

import (
	"database/sql"
	"gqlgen/cmd/graph/model"
)

type TodoModel struct {
	DB *sql.Conn
}

func (t *TodoModel) GetTodo(id string) (*model.NewTodo, error) {

	return nil
}

func (t *TodoModel) MakeTodo(todoInput *model.NewTodo) error {

	return nil
}
