package result

import (
	"testing"
)

func TestResult_IsSuccessful(t *testing.T) {
	// Test successful result
	result := Ok()
	if !result.IsSuccessful() {
		t.Error("Expected new result to be successful")
	}

	// Test failed result
	failedResult := Fail(NewError("test error"))
	if failedResult.IsSuccessful() {
		t.Error("Expected result with error to be unsuccessful")
	}
}

func TestResult_IsFailed(t *testing.T) {
	// Test successful result
	result := Ok()
	if result.IsFailed() {
		t.Error("Expected new result to not be failed")
	}

	// Test failed result
	failedResult := Fail(NewError("test error"))
	if !failedResult.IsFailed() {
		t.Error("Expected result with error to be failed")
	}
}

func TestResult_Fail(t *testing.T) {
	// Test with valid error
	err := NewError("test error")
	result := Fail(err)

	if len(result.GetErrors()) != 1 {
		t.Errorf("Expected 1 error, got %d", len(result.GetErrors()))
	}

	// Test with nil error (should create default error)
	result = Fail(nil)
	if len(result.GetErrors()) != 1 {
		t.Errorf("Expected 1 default error, got %d", len(result.GetErrors()))
	}

	// Test with multiple errors
	result = Fail(NewError("error1"), NewError("error2"))
	if len(result.GetErrors()) != 2 {
		t.Errorf("Expected 2 errors, got %d", len(result.GetErrors()))
	}
}

func TestResult_AddError(t *testing.T) {
	result := Ok()
	err := NewError("test error")

	result.AddError(err)
	if len(result.GetErrors()) != 1 {
		t.Errorf("Expected 1 error, got %d", len(result.GetErrors()))
	}

	// Test adding nil error (should not add)
	result.AddError(nil)
	if len(result.GetErrors()) != 1 {
		t.Errorf("Expected still 1 error after adding nil, got %d", len(result.GetErrors()))
	}
}

func TestResult_AddRange(t *testing.T) {
	result := Ok()

	// Add multiple items
	result.AddRange(
		NewError("error1"),
		NewError("error2"),
	)

	if len(result.GetErrors()) != 2 {
		t.Errorf("Expected 2 errors, got %d", len(result.GetErrors()))
	}

	// Test with nil values (should be ignored)
	result = Ok()
	result.AddRange(nil, NewError("error"), nil)

	if len(result.GetErrors()) != 1 {
		t.Errorf("Expected 1 error, got %d", len(result.GetErrors()))
	}

}

func TestError_GetMessage(t *testing.T) {
	message := "test error message"
	err := NewError(message)

	if err.GetMessage() != message {
		t.Errorf("Expected message '%s', got '%s'", message, err.GetMessage())
	}
}
