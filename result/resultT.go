package result

// ResultT represents a result that can contain errors and warnings
type ResultT[T any] struct {
	Result
	value T
}

// NewResultT creates a new ResultT instance
func NewResultT[T any](value T) *ResultT[T] {
	return &ResultT[T]{
		Result: *NewResult(),
		value:  value,
	}
}

// GetPayload returns the payload
func (r *ResultT[T]) GetValue() T {
	return r.value
}

// OkT creates a new successful ResultT with the given payload
func OkT[T any](value T) *ResultT[T] {
	return NewResultT(value)
}

// // FailT creates a new failed ResultT with the given errors
// func FailT[T any](errors ...*Error) *ResultT[T] {
// 	var zero T
// 	result := NewResultT(zero)

// 	// Filter out nil errors
// 	validErrors := make([]*Error, 0)
// 	for _, err := range errors {
// 		if err != nil {
// 			validErrors = append(validErrors, err)
// 		}
// 	}

// 	// If no valid errors provided, add a default error
// 	if len(validErrors) == 0 {
// 		validErrors = append(validErrors, NewError("Fail result is created without any error."))
// 	}

// 	result.errors = validErrors
// 	return result
// }
