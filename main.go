package main

import (
	"bitbucket.org/Aleksqew/test-task/api-gateway/application/factories"
	"bitbucket.org/Aleksqew/test-task/api-gateway/infrastructure/configs"
	"bitbucket.org/Aleksqew/test-task/api-gateway/infrastructure/repository"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	globalConfigs := configs.GetGlobalConfigs()
	rep := repository.NewTodoRepository()
	serviceApi := factories.NewServiceApi(*rep)
	api := factories.NewControllerApi(globalConfigs.GetEchoConfig(), serviceApi)
	api.Init()
}
