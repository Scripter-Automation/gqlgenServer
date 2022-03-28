package model2

import

type Todo interface{
	GetTodo(id string) (*model.NewTodo, error)
	MakeTodo(todoInput *model.NewTodo) error
}