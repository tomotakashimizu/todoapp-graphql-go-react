// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateTodoInput struct {
	Text string `json:"text"`
}

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"userId"`
	User   *User  `json:"user"`
}

type UpdateTodoInput struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
