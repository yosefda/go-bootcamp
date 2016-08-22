// This is the main package which is the entry point to our
// program.
package main

import (
	// This is go standard lib package.
	"fmt"

	// This is our custom lib package.
	"github.com/yosefda/stringutil"
)

// main package must have func main.
func main() {
	reversedStr := "!oG ,olleH"
	fmt.Printf("%s --> %s\n", reversedStr, stringutil.Reverse(reversedStr))
}
