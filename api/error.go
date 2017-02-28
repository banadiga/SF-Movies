package api

type Error struct {
	Message string  `json:"message" binding:"required"`
}

func NewError(message error) (*Error) {
	return &Error{
		Message: message.Error(),
	}
}
