package factories

import (
	"bitbucket.org/Aleksqew/test-task/api-gateway/application/services"
	"bitbucket.org/Aleksqew/test-task/api-gateway/infrastructure/repository"
)

func NewServiceApi(repo repository.TodoRepo) services.ServiceApi {
	api := services.ServiceApi{}
	api.SetTodoRepository(&repo)
	return api
}
