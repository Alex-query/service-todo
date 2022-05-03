package factories

import (
	"bitbucket.org/Aleksqew/test-task/api-gateway/application/controller"
	"bitbucket.org/Aleksqew/test-task/api-gateway/application/services"
	"bitbucket.org/Aleksqew/test-task/api-gateway/infrastructure/configs"
)

func NewControllerApi(config configs.EchoConfig, serviceApi services.ServiceApi) controller.Api {
	api := controller.Api{}
	api.SetEchoConfig(config)
	api.SetServiceApi(serviceApi)
	return api
}