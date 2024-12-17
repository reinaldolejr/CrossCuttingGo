package result

import (
	"testing"
)

func TestResultT(t *testing.T) {
	// Test successful result
	successResult := OkT("test payload")
	if !successResult.IsSuccessful() {
		t.Error("Expected result to be successful")
	}
	if successResult.GetValue() != "test payload" {
		t.Error("Expected payload to be 'test payload'")
	}

}
