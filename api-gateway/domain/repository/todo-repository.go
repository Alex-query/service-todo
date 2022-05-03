package repository

import "bitbucket.org/Aleksqew/test-task/api-gateway/domain/entity"

type TodoRepository interface {
	CreateTodo(todo entity.Todo) (*entity.Todo, error)
	EditTodo(todo entity.Todo) (*entity.Todo, error)
	GetTodo(id string) (*entity.Todo, error)
	ListTodo(ownerId int64) ([]entity.Todo, error)
}
