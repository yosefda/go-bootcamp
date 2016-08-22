package main

import (
	"fmt"
)

// This is an unexported const. Unexported consts are not
// available outside the package in which they were declared.
// The naming starts with lowercase and they dont need
// comment.
const pi = 3.14

// These are exported consts. Exported consts are avilable
// outside the package in which they were declared. The
// naming is camelcase and they need to be commented.
const (
	IdeKey = "Atom"
	Yes    = true
	No     = false
)

func main() {
	fmt.Printf("Pi is %f\n", pi)
	fmt.Printf("IDE is %s\n", IdeKey)
	fmt.Printf("YES is %t and NO is %t\n", Yes, No)
}
