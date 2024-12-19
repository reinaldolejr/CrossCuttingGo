# CrossCuttingGo

CrossCuttingGo is a lightweight Go library that provides a Result type for better error and warning handling in Go applications. It helps you manage operation outcomes with success/failure states, and errors in a more structured way.

## Features

- 🎯 Simple Result type with success/failure states
- ⚠️ Support for multiple errors
- 🔄 Fluent interface for chaining operations
- 🔌 Easy integration with existing Go projects
- 🧪 Type-safe generic Result[T] for handling typed returns

## Installation

```bash
go get github.com/reinaldolejr/CrossCuttingGo
```

## Quick Start

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/reinaldolejr/CrossCuttingGo/result"
)

func main() {
    // Create a successful result
    successResult := result.Ok()
    fmt.Printf("Success: %v\n", successResult.IsSuccessful())

    // Create a failed result
    failedResult := result.Fail(result.NewError("Operation failed"))
    fmt.Printf("Failed: %v\n", failedResult.IsFailed())

    // Chain multiple operations
    r.AddRange(
        result.NewError("Something went wrong"),
        result.NewError("Connection timeout"),
    )
}
```

### Using Generic Results

```go
package main

import (
    "fmt"
    "github.com/reinaldolejr/CrossCuttingGo/result"
)

func divide(a, b float64) result.Result[float64] {
    if b == 0 {
        return result.FailT[float64](result.NewError("Division by zero"))
    }
    return result.OkT(a / b)
}

func main() {
    r := divide(10, 2)
    if r.IsSuccessful() {
        value := r.Value() // 5
        fmt.Printf("Result: %v\n", value)
    }

    r2 := divide(10, 0)
    if r2.IsFailed() {
        fmt.Printf("Error: %v\n", r2.Errors()[0].Message())
    }
}
```

## API Reference

### Result Types

- `Result` - Basic result type for success/failure scenarios
- `Result[T]` - Generic result type that includes a value of type T

### Creating Results

```go
// Basic Results
result.Ok()                    // Successful result
result.Fail(error)            // Failed result with error

// Generic Results
result.OkT(value T)           // Successful result with value
result.FailT[T](error)        // Failed result with error
```

### Methods

#### Common Methods
- `IsSuccessful() bool` - Checks if the result is successful
- `IsFailed() bool` - Checks if the result has failed
- `AddError(error)` - Adds an error
- `AddRange(items...)` - Adds multiple errors
- `Errors() []Error` - Gets all errors

#### Generic Result Methods
- `Value() T` - Gets the value (for Result[T])
- `ValueOrDefault() T` - Gets the value or zero value if failed

## Error Handling Best Practices

```go
func ProcessOrder(orderID string) result.Result[Order] {
    order, err := db.GetOrder(orderID)
    if err != nil {
        return result.FailT[Order](result.NewError("Failed to fetch order"))
    }

    if order.IsExpired() {
        r := result.OkT(order)
        r.AddWarning(result.NewWarning("Order is expired"))
        return r
    }

    return result.OkT(order)
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

Reinaldo Junior ([@reinaldolejr](https://github.com/reinaldolejr))