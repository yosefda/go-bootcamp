package main

import "fmt"

// Point Point type
type Point struct {
	X, Y int
}

var (
	s = Point{1, 2}  // has type Point
	p = &Point{3, 4} // has type *Point
	u = Point{X: 1}  // Y:0 is implicit
	v = Point{}      // X:0 and Y:0
)

func main() {
	fmt.Println(s, p, u, v)
}
