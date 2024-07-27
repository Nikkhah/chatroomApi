package request

type CreateRoomRequest struct {
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

type JoinRoomRequest struct {
	UserID     string `json:"user_id"`
	ChatroomID string `json:"chatroom_id"`
}
