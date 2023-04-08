package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/tomotakashimizu/todoapp-graphql-go-react/graph/model"
	"github.com/tomotakashimizu/todoapp-graphql-go-react/infrastructure"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.InputTodo) (bool, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		log.Printf("Error generating UUID: %v\n", err)
		return false, err
	}
	todo := &infrastructure.Todo{
		ID:     newUUID.String(),
		Text:   input.Text,
		UserID: input.UserID,
	}

	_, err = r.DB.Insert(todo)
	if err != nil {
		log.Printf("Error insert todo: %v\n", err)
		return false, err
	}

	return true, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*infrastructure.Todo
	r.DB.AllCols().Find(&todos)
	return model.NewTodos(todos), nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Todo returns TodoResolver implementation.
func (r *Resolver) Todo() TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
