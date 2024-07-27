package service

import (
	"chatroom/internal/core/model/request"
	"chatroom/internal/core/model/response"
)

type ChatroomService interface {
	Create(request *request.CreateRoomRequest) *response.Response
	Join(request *request.JoinRoomRequest) *response.Response
}
