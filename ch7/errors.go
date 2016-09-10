package main

import (
	"fmt"
	"time"
)

// MyError Custom error type
type MyError struct {
	When time.Time
	What string
}

// Error Print error message (implicitly implement error
// interface)
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// run Return error type. This demonstrates that MyError
// is of error type
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
