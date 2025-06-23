package main

import (
    "errors"
    "fmt"
)

// CustomError defines a custom error type with a code and message.

type CustomError struct {
    Code    int
    Message string
}

func checkNumber(num int) (string, error) {
    if num < 0 {
        return "", errors.New("number is negative")
    }
    return "number is positive", nil
}

// Error implements the error interface for CustomError.

func (e *CustomError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// generateError creates and returns a new CustomError.
func generateError() error {
    return &CustomError{Code: 404, Message: "Resource not found"}
}



func main() {
    result, err := checkNumber(-5)
	
    if err != nil {
        fmt.Println("Error:", err) // Output: Error: number is negative
    } else {
        fmt.Println(result)
    }
	  err2 := generateError()
    fmt.Println(err2) 
}

//When there’s no error, return nil to indicate success.
// return "operation successful", nil

//wrapping errors
// better error tracing, Go provides fmt.Errorf to wrap errors with additional context.
// Unwrapping Errors
// To extract the original error from a wrapped error, use errors.Unwrap.




