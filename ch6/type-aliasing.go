package main

import (
	"fmt"
	"strings"
)

// MyString Type aliasing for string type
type MyString string

// Uppercase Convert MyString to uppercase
func (ms MyString) Uppercase() string {
	// need to cast MyString to string again
	// as ToUpper method only accepts string type
	return strings.ToUpper(string(ms))
}

func main() {
	str := MyString("This is MyString")
	fmt.Println(str.Uppercase())
}
