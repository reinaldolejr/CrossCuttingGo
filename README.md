# CrossCuttingGo

CrossCuttingGo is a lightweight Go library that provides a Result type for better error and warning handling in Go applications. It helps you manage operation outcomes with success/failure states, errors, and warnings in a more structured way.

## Features

- Simple Result type with success/failure states
- Support for multiple errors and warnings
- Fluent interface for adding errors and warnings
- Easy to integrate into existing Go projects

## Installation

To install CrossCuttingGo, use `go get`:

```bash
go get github.com/reinaldolejr/CrossCuttingGo
```

## Usage

Here's a simple example of how to use CrossCuttingGo:

```go

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
    // Add warnings to a result
    successResult.AddWarning(result.NewWarning("This is a warning"))
    // Add multiple items (errors and warnings)
    successResult.AddRange(
        result.NewError("An error"),
        result.NewWarning("Another warning"),
    )
}

```


## API Reference

### Creating Results

- `result.Ok()` - Creates a new successful result
- `result.Fail(error)` - Creates a new failed result with an error

### Methods

- `IsSuccessful()` - Returns true if the result is successful
- `IsFailed()` - Returns true if the result has failed
- `AddWarning(warning)` - Adds a warning to the result
- `AddError(error)` - Adds an error to the result
- `AddRange(items...)` - Adds multiple errors and/or warnings to the result

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the [LICENSE](LICENSE) file in the root directory of this repository.

## Author

Reinaldo Junior (@reinaldolejr)