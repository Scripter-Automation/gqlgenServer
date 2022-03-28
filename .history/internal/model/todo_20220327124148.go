package model

import (
	"gqlgen/cmd/graph/model"
)

type Todo interface {
	GetTodo(id string) (*model.NewTodo, error)
	MakeTodo(todoInput *model.NewTodo) error
}
