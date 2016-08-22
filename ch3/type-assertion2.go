package main

import "fmt"

// Stringer Stringer interface.
type Stringer interface {
	String() string
}

type fakeString struct {
	content string
}

// fakeString implements Stringer.
func (s *fakeString) String() string {
	return s.content
}

// This function perform type assertion on its param.
func printString(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Println(value)
	case Stringer:
		fmt.Println(str.String())
	}
}

func main() {
	fs := &fakeString{"This is a fake string"}
	printString(fs)
	printString("This is a builtin string")
}
