package controller

import (
	"bitbucket.org/Aleksqew/test-task/api-gateway/application/requests"
	"bitbucket.org/Aleksqew/test-task/api-gateway/application/responses"
	"bitbucket.org/Aleksqew/test-task/api-gateway/application/services"
	"bitbucket.org/Aleksqew/test-task/api-gateway/infrastructure/configs"
	"net/http"
	"github.com/go-playground/validator"
)
import "github.com/labstack/echo"

type Api struct {
	echo       *echo.Echo
	echoConfig configs.EchoConfig
	serviceApi services.ServiceApi
}

func (api *Api) Init() {
	api.echo = echo.New()
	api.RegisterRouteBase()
	api.echo.Validator = &CustomValidator{validator: validator.New()}
	api.echo.Start(":" + api.echoConfig.GetPort())
}

func (api *Api) RegisterRouteBase() {
	api.echo.GET("/api/ping", api.ping)
	api.echo.POST("/api/todo", api.todoCreate)
	api.echo.GET("/api/todo/:id", api.todoGet)
}

func (api *Api) ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "pong")
}

func (api *Api) todoCreate(c echo.Context) error {
	requestCreateTodo := new(requests.CreateTodoDto)
	if err := c.Bind(requestCreateTodo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, api.generateErrorResponse(err))
	}
	if err := c.Validate(requestCreateTodo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, api.generateErrorResponse(err))
	}
	response, err := api.serviceApi.CreateTodo(*requestCreateTodo)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, api.generateSuccessResponse(response))
}

func (api *Api) todoGet(c echo.Context) error {
	response, err := api.serviceApi.GetTodo(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, api.generateErrorResponse(err))
	}
	return c.JSON(http.StatusOK, api.generateSuccessResponse(response))
}

func (api *Api) generateSuccessResponse(response interface{}) responses.BaseDto {
	b := responses.BaseDto{
		Data: response,
		Meta: responses.Meta{
			Status: responses.META_STATUS_SUCCESS,
		},
	}
	return b
}

func (api *Api) generateErrorResponse(err error) responses.BaseDto {
	b := responses.BaseDto{
		Data: err.Error(),
		Meta: responses.Meta{
			Status: responses.META_STATUS_ERROR,
		},
	}
	return b
}

func (api *Api) SetEchoConfig(config configs.EchoConfig) {
	api.echoConfig = config
}

func (api *Api) SetServiceApi(serviceApi services.ServiceApi) {
	api.serviceApi = serviceApi
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
