package service

import (
	"chatroom/internal/core/common/utils"
	"chatroom/internal/core/dto"
	"chatroom/internal/core/entity/error_code"
	"chatroom/internal/core/model/request"
	"chatroom/internal/core/model/response"
	"chatroom/internal/core/port/repository"
	"chatroom/internal/core/port/service"
	"fmt"
	"strconv"
)

type chatroomService struct {
	chatroomRepo repository.ChatroomRepository
}

const (
	invalidNameErrMsg = "invalid name"
	invalidIDErrMsg   = "invalid id"
)

func NewChatroomService(chatroomRepo repository.ChatroomRepository) service.ChatroomService {
	return &chatroomService{
		chatroomRepo: chatroomRepo,
	}
}

func (ch chatroomService) Join(request *request.JoinRoomRequest) *response.Response {
	//validate request
	if len(request.ChatroomID) == 0 || len(request.UserID) == 0 {
		return ch.failedResponse(error_code.InvalidRequest, invalidNameErrMsg)
	}
	userID, err := strconv.Atoi(request.UserID)
	if err != nil {
		return ch.failedResponse(error_code.InternalError, error_code.InternalErrMsg)
	}
	chatroomID, err := strconv.Atoi(request.ChatroomID)
	if err != nil {
		return ch.failedResponse(error_code.InternalError, error_code.InternalErrMsg)
	}
	userChatroomDTO := dto.UserChatroomDTO{
		UserID:     userID,
		ChatroomID: chatroomID,
	}
	// save a new user
	{
		err := ch.chatroomRepo.InsertUserChatroom(userChatroomDTO)
		if err != nil {
			fmt.Println(err.Error())
			if err == repository.DuplicateChatroom {
				return ch.failedResponse(error_code.DuplicateChatroom, err.Error())
			}
			return ch.failedResponse(error_code.InternalError, error_code.InternalErrMsg)
		}
	}

	// create data response
	createData := response.CreateDataResponse{}
	return ch.successResponse(createData)
}

func (ch chatroomService) Create(request *request.CreateRoomRequest) *response.Response {
	// validate request
	if len(request.Name) == 0 {
		return ch.failedResponse(error_code.InvalidRequest, invalidNameErrMsg)
	}

	currentTime := utils.GetUTCCurrentMillis()
	userID, err := strconv.Atoi(request.UserID)
	if err != nil {
		return ch.failedResponse(error_code.InternalError, error_code.InternalErrMsg)
	}
	fmt.Println(">>>> user id ", userID)
	chatroomDTO := dto.ChatroomDTO{
		Name:      request.Name,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID:    userID,
	}

	// save a new chatroom
	insertedID, err := ch.chatroomRepo.InsertChatroom(chatroomDTO)
	if err != nil {
		fmt.Println(err.Error())
		if err == repository.DuplicateChatroom {
			return ch.failedResponse(error_code.DuplicateChatroom, err.Error())
		}
		return ch.failedResponse(error_code.InternalError, error_code.InternalErrMsg)
	}

	// create data response
	createData := response.CreateDataResponse{
		Name: chatroomDTO.Name,
		ID:   insertedID,
	}
	return ch.successResponse(createData)
}

func (ch chatroomService) failedResponse(
	code error_code.ErrorCode, message string,
) *response.Response {
	return &response.Response{
		Status:       false,
		ErrorCode:    code,
		ErrorMessage: message,
	}
}

func (ch chatroomService) successResponse(data response.CreateDataResponse) *response.Response {
	return &response.Response{
		Data:         data,
		Status:       true,
		ErrorCode:    error_code.Success,
		ErrorMessage: error_code.SuccessErrMsg,
	}
}
