package repository

import (
	"chatroom/internal/core/dto"
	"errors"
)

var (
	DuplicateChatroom = errors.New("duplicate chatroom")
)

type ChatroomRepository interface {
	InsertChatroom(chatroom dto.ChatroomDTO) (string, error)
	InsertUserChatroom(userChatroom dto.UserChatroomDTO) error
}
