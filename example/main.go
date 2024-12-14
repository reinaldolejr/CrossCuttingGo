package main

import (
	"fmt"

	"github.com/reinaldolejr/CrossCuttingGo/result"
)

func main() {
	// Create a successful result
	successResult := result.Ok()
	fmt.Printf("Is successful: %v\n", successResult.IsSuccessful())

	// Create a failed result
	failedResult := result.Fail(result.NewError("Something went wrong"))
	fmt.Printf("Is failed: %v\n", failedResult.IsFailed())

	// Add warnings
	successResult.AddWarning(result.NewWarning("This is a warning"))

	// Add multiple items
	successResult.AddRange(
		result.NewError("An error"),
		result.NewWarning("Another warning"),
	)
}
