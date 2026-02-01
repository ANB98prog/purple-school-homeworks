package errors

import "errors"

// ItemNotFound - ошибка возвращаемая при ненайденной сущности
type ItemNotFound struct {
	Message string `json:"message"`
}

func (item *ItemNotFound) Error() string {
	return item.Message
}

var (
	ErrUserNotFound = errors.New("user not found")
)
