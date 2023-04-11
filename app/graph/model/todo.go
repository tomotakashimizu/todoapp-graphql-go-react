package model

import "github.com/tomotakashimizu/todoapp-graphql-go-react/app/infrastructure"

func NewTodos(todos []*infrastructure.Todo) []*Todo {
	data := make([]*Todo, len(todos))
	for i, v := range todos {
		data[i] = NewTodo(v)
	}
	return data
}

func NewTodo(todo *infrastructure.Todo) *Todo {
	return &Todo{
		ID:     todo.ID,
		Text:   todo.Text,
		Done:   todo.Done,
		UserID: todo.UserID,
		User:   &User{ID: todo.UserID, Name: "user " + todo.UserID},
	}
}
