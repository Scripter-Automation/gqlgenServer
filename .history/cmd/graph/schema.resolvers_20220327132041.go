package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gqlgen/cmd/graph/generated"
	"gqlgen/cmd/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (bool, error) {
	err := r.TodoModel.MakeTodo(input)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *queryResolver) Todo(ctx context.Context, input string) (*model.Todo, error) {
	todo, err := r.TodoModel.GetTodo(input)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
