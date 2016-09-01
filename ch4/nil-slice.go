package main

import "fmt"

func main() {
	var s []string

	// len() checks length of a slice
	// cap() checks capacity of a slice
	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil")
	}
}
