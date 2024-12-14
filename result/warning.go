package result

// Warning represents a warning in the result
type Warning struct {
	Message string
}

func NewWarning(message string) *Warning {
	return &Warning{Message: message}
}

func (w *Warning) GetMessage() string {
	return w.Message
}
