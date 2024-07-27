package response

import "chatroom/internal/core/entity/error_code"

type Response struct {
	Data         interface{}          `json:"data"`
	Status       bool                 `json:"status"`
	ErrorCode    error_code.ErrorCode `json:"errorCode"`
	ErrorMessage string               `json:"errorMessage"`
}

type CreateDataResponse struct {
	Name string `json:"Name"`
	ID   string `json:"id"`
}
