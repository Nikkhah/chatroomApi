package dto

type ChatroomDTO struct {
	Name      string
	CreatedAt uint64
	UpdatedAt uint64
	UserID    int
}

type UserChatroomDTO struct {
	UserID     int
	ChatroomID int
}
