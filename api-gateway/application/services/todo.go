package services

import (
	"bitbucket.org/Aleksqew/test-task/api-gateway/application/requests"
	"bitbucket.org/Aleksqew/test-task/api-gateway/application/responses"
	"bitbucket.org/Aleksqew/test-task/api-gateway/domain/entity"
	"bitbucket.org/Aleksqew/test-task/api-gateway/domain/repository"
)

type ServiceApi struct {
	todoRepository repository.TodoRepository
}

func (serviceApi *ServiceApi) Init() {

}

func (serviceApi *ServiceApi) CreateTodo(requestDto requests.CreateTodoDto) (responses.TodoDto, error) {
	todo, err := serviceApi.todoRepository.CreateTodo(entity.Todo{
		Title:   requestDto.Title,
		OwnerId: 0,
	})
	if err != nil {
		return responses.TodoDto{}, err
	}
	resp:=responses.TodoDto{
		Id:    todo.Id,
		Title: todo.Title,
	}
	return resp, nil
}

func (serviceApi *ServiceApi) GetTodo(id string) (responses.TodoDto, error) {
	todo, err := serviceApi.todoRepository.GetTodo(id)
	if err != nil {
		return responses.TodoDto{}, err
	}
	resp:=responses.TodoDto{
		Id:    todo.Id,
		Title: todo.Title,
	}
	return resp, nil
}

func (serviceApi *ServiceApi) SetTodoRepository(repository repository.TodoRepository) {
	serviceApi.todoRepository = repository
}
