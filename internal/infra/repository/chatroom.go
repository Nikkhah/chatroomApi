package repository

import (
	"chatroom/internal/core/dto"
	"chatroom/internal/core/port/repository"
	"errors"
	"fmt"
	"strings"
)

const (
	duplicateEntryMsg = "Duplicate entry"
	numberRowInserted = 1
)

var (
	insertChatroomErr = errors.New("failed to insert chatroom")
)

const (
	insertChatroomStatement = "INSERT INTO Chatroom ( " +
		"`name`, " +
		"`user_id`," +
		"`created_at`," +
		"`updated_at`) " +
		"VALUES (?, ?, ?,?)"

	insertUserChatroomStatement = "INSERT INTO User_Chatroom ( " +
		"`user_id`," +
		"`chatroom_id`) " +
		"VALUES (?, ?)"
)

type chatroomRepository struct {
	db repository.Database
}

func NewChatroomRepository(db repository.Database) repository.ChatroomRepository {
	return &chatroomRepository{
		db: db,
	}
}

func (ch chatroomRepository) InsertUserChatroom(userRoom dto.UserChatroomDTO) error {
	_, err := ch.db.GetDB().Exec(
		insertUserChatroomStatement,
		userRoom.UserID,
		userRoom.ChatroomID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (ch chatroomRepository) InsertChatroom(room dto.ChatroomDTO) (string, error) {
	result, err := ch.db.GetDB().Exec(
		insertChatroomStatement,
		room.Name,
		room.CreatedAt,
		room.UpdatedAt,
		room.UserID,
	)

	if err != nil {
		fmt.Println(" >> ", err.Error())
		if strings.Contains(err.Error(), duplicateEntryMsg) {
			return "", repository.DuplicateChatroom
		}
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", id), nil
}
