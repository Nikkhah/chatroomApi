package error_code

type ErrorCode string

// error code
const (
	Success           ErrorCode = "SUCCESS"
	InvalidRequest    ErrorCode = "INVALID_REQUEST"
	DuplicateChatroom ErrorCode = "DUPLICATE_CHATROOM"
	InternalError     ErrorCode = "INTERNAL_ERROR"
)

// error message
const (
	SuccessErrMsg        = "success"
	InternalErrMsg       = "internal error"
	InvalidRequestErrMsg = "invalid request"
)
