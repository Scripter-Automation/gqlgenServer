package db

import (
	"database/sql"
	"gqlgen/cmd/graph/model"

	"github.com/dgryski/trifles/uuid"
)

type TodoModel struct {
	DB *sql.DB
}

func (t *TodoModel) GetTodo(id string) (*model.Todo, error) {
	var todo model.Todo
	query := "SELECT id, text FROM todo"
	row := t.DB.QueryRow(query, id)

	err := row.Scan(&todo.ID, &todo.Text)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}
func (t *TodoModel) MakeTodo(todoInput *model.NewTodo) error {
	id := uuid.UUIDv4()
	query := "INSERT INTO todo (id, text) values (?, ?)"
	_, err := t.DB.Exec(query, id, todoInput.Text)
	if err != nil {
		return err
	}
	return nil
}
