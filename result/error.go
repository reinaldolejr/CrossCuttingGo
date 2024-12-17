package result

// Error represents an error in the result
type Error struct {
	Message string
}

// NewError creates a new Error instance
func NewError(message string) *Error {
	return &Error{Message: message}
}

// GetMessage returns the error message
func (e *Error) GetMessage() string {
	return e.Message
}
