package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.29

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/tomotakashimizu/todoapp-graphql-go-react/app/graph/model"
	"github.com/tomotakashimizu/todoapp-graphql-go-react/app/infrastructure"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.CreateTodoInput) (*model.Todo, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		log.Printf("Error generating UUID: %v\n", err)
		return nil, err
	}
	todo := &infrastructure.Todo{
		ID:   newUUID.String(),
		Text: input.Text,
		// TODO: ログインしているユーザのIDにする
		UserID: "550e8400-e29b-41d4-a716-446655440000",
	}

	_, err = r.DB.Insert(todo)
	if err != nil {
		log.Printf("Error insert todo: %v\n", err)
		return nil, err
	}

	return model.NewTodo(todo), nil
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, todoID string, input model.UpdateTodoInput) (bool, error) {
	todo := &infrastructure.Todo{
		Text: input.Text,
		Done: input.Done,
	}
	_, err := r.DB.ID(todoID).Cols("text", "done").Update(todo)
	if err != nil {
		log.Printf("Error update todo: %v\n", err)
		return false, err
	}
	return true, nil
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, todoID string) (bool, error) {
	_, err := r.DB.ID(todoID).Delete(&infrastructure.Todo{})
	if err != nil {
		log.Printf("Error delete todo: %v\n", err)
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
