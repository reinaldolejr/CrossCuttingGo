package result

// Error represents an error in the result
type Error struct {
	Message string
}

func NewError(message string) *Error {
	return &Error{Message: message}
}

func (e *Error) GetMessage() string {
	return e.Message
}
