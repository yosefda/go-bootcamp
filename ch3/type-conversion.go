package main

import (
	"fmt"
)

func main() {
	var i = 42
	var f = float64(i)
	var u = uint(f)

	fmt.Printf("i => %T (%v)\n", i, i)
	fmt.Printf("float(i) => %T (%v)\n", f, f)
	fmt.Printf("uint(f) => %T (%v)\n", u, u)
}
