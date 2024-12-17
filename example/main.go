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

	failedList := []result.Error{
		*result.NewError("Something went wrong"),
		*result.NewError("Something went wrong 2"),
	}
	fmt.Printf("Failed list: %v\n", failedList)

	successResultT := result.OkT[any](struct {
		Id   int
		Name string
	}{
		Id:   1,
		Name: "John",
	})
	fmt.Printf("Is successful T: %v\n", successResultT.IsSuccessful())
	fmt.Printf("Value T: %v\n", successResultT.GetValue())

	// Add multiple items
	successResult.AddRange(
		result.NewError("An error"),
	)
}
