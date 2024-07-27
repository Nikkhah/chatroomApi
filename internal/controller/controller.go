package controller

import (
	"chatroom/internal/core/common/router"
	"chatroom/internal/core/entity/error_code"
	"chatroom/internal/core/model/request"
	"chatroom/internal/core/model/response"
	"chatroom/internal/core/port/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	invalidRequestResponse = &response.Response{
		ErrorCode:    error_code.InvalidRequest,
		ErrorMessage: error_code.InvalidRequestErrMsg,
		Status:       false,
	}
)

type ChatroomController struct {
	gin             *gin.Engine
	chatroomService service.ChatroomService
}

func NewChatroomController(
	gin *gin.Engine,
	chatroomService service.ChatroomService,
) ChatroomController {
	return ChatroomController{
		gin:             gin,
		chatroomService: chatroomService,
	}
}

func (ch ChatroomController) InitRouter() {
	api := ch.gin.Group("/api/v1")
	router.Post(api, "/create/chatroom", ch.create)
	router.Post(api, "/join/chatroom", ch.join)
}

func (ch ChatroomController) create(c *gin.Context) {
	req, err := ch.parseChatroomRequest(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &invalidRequestResponse)
		return
	}
	resp := ch.chatroomService.Create(req)
	c.JSON(http.StatusOK, resp)
}

func (ch ChatroomController) join(c *gin.Context) {
	req, err := ch.parseUserChatroomRequest(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, &invalidRequestResponse)
		return
	}
	resp := ch.chatroomService.Join(req)
	c.JSON(http.StatusOK, resp)
}

func (ch ChatroomController) parseChatroomRequest(ctx *gin.Context) (*request.CreateRoomRequest, error) {

	var req request.CreateRoomRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &req, nil
}

func (ch ChatroomController) parseUserChatroomRequest(ctx *gin.Context) (*request.JoinRoomRequest, error) {
	var req request.JoinRoomRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &req, nil
}
