package responses

type BaseDto struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

const META_STATUS_SUCCESS = "success"
const META_STATUS_ERROR = "error"
const META_STATUS_VALIDATE_ERROR = "validate_error"
const META_STATUS_AUTH_ERROR = "auth_error"
const META_STATUS_PERMISSION_ERROR = "permission_error"

type Meta struct {
	Status string `json:"status"`
}
