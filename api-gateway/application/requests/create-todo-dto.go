package requests

type CreateTodoDto struct {
	Title string `json:"title" validate:"required"`
}
