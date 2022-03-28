package db

import (
	"database/sql"
	"gqlgen/cmd/graph/model"
)

type TodoModel struct {
	DB *sql.DB
}

func (t *TodoModel) GetTodo(id string) (*model.Todo, error) {

	var todo model.Todo

	query:= "SELECT id, text FROM todo"
	row:=	t.DB.QueryRow(query, id)

	err:= row.Scan(&todo.ID, &todo.Text)
	if err!=nil{
		return err
	}

	return &, nil
}

func (t *TodoModel) MakeTodo(todoInput *model.NewTodo) error {

	return nil
}
