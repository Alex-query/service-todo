package repository

import (
	"bitbucket.org/Aleksqew/test-task/api-gateway/domain/entity"
	"bitbucket.org/Aleksqew/test-task/api-gateway/domain/repository"
	"errors"
	"github.com/google/uuid"
)

type TodoRepo struct {
	storageTodo []entity.Todo
}

func NewTodoRepository() *TodoRepo {
	return &TodoRepo{}
}

var _ repository.TodoRepository = &TodoRepo{}

func (t *TodoRepo) Init() {
	t.storageTodo = []entity.Todo{}
}

func (t *TodoRepo) CreateTodo(todo entity.Todo) (*entity.Todo, error) {
	id := uuid.New()
	todo.Id = id.String()
	t.storageTodo = append(t.storageTodo, todo)
	return &todo, nil
}

func (t *TodoRepo) EditTodo(todoIn entity.Todo) (*entity.Todo, error) {
	todo, err := t.GetTodo(todoIn.Id)
	if err != nil {
		return nil, err
	}
	todo.Title = todoIn.Title
	return todo, nil
}

func (t *TodoRepo) GetTodo(id string) (*entity.Todo, error) {
	for _, v := range t.storageTodo {
		if v.Id == id {
			return &v, nil
		}
	}
	return nil, errors.New("not found")
}

func (t *TodoRepo) ListTodo(ownerId int64) ([]entity.Todo, error) {
	return t.storageTodo, nil
}
