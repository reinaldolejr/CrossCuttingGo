package result

// Result represents a result that can contain errors and warnings
type Result struct {
	errors   []*Error
	warnings []*Warning
}

// NewResult creates a new Result instance
func NewResult() *Result {
	return &Result{
		errors:   make([]*Error, 0),
		warnings: make([]*Warning, 0),
	}
}

// IsSuccessful returns true if the result has no errors
func (r *Result) IsSuccessful() bool {
	return len(r.errors) == 0
}

// IsFailed returns true if the result has errors
func (r *Result) IsFailed() bool {
	return !r.IsSuccessful()
}

// GetErrors returns the slice of errors
func (r *Result) GetErrors() []*Error {
	return r.errors
}

// GetWarnings returns the slice of warnings
func (r *Result) GetWarnings() []*Warning {
	return r.warnings
}

// Ok creates a new successful Result
func Ok() *Result {
	return NewResult()
}

// Fail creates a new failed Result with the given errors
func Fail(errors ...*Error) *Result {
	result := NewResult()

	// Filter out nil errors
	validErrors := make([]*Error, 0)
	for _, err := range errors {
		if err != nil {
			validErrors = append(validErrors, err)
		}
	}

	// If no valid errors provided, add a default error
	if len(validErrors) == 0 {
		validErrors = append(validErrors, NewError("Fail result is created without any error."))
	}

	result.errors = validErrors
	return result
}

// AddError adds an error to the result
func (r *Result) AddError(err *Error) {
	if err != nil {
		r.errors = append(r.errors, err)
	}
}

// AddWarning adds a warning to the result
func (r *Result) AddWarning(warning *Warning) {
	if warning != nil {
		r.warnings = append(r.warnings, warning)
	}
}

// AddRange adds multiple errors or warnings to the result
func (r *Result) AddRange(items ...interface{}) {
	for _, item := range items {
		switch v := item.(type) {
		case *Error:
			r.AddError(v)
		case *Warning:
			r.AddWarning(v)
		}
	}
}
