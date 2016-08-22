package main

import "fmt"

// Bootcamp Bootcamp type
type Bootcamp struct {
	Lat, Lon float64
}

func main() {
	// The initialisation function new() will allocated memory
	// for the variable. The content of the allocated memory is
	// zeroed out, thus it creates zero value of the type.
	x := new(int)
	fmt.Printf("x := new(int) contains type: %T and value: %v -> %v\n", x, x, *x)

	s := new(Bootcamp)
	p := &Bootcamp{}
	fmt.Printf("%T --> %+v\n", s, s)
	fmt.Printf("%T --> %+v\n", p, p)
}
