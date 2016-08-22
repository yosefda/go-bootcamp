package main

import (
	"fmt"
	"math/cmplx"
)

// One way of declaring multiple variables.
var (
	goIsFun        = true                 // type will be infered
	maxInt  uint64 = 1<<64 - 1            // type must be specified
	complex        = cmplx.Sqrt(-5 + 12i) // type will be infered
)

func main() {
	const f = "%T (%v)\n"
	fmt.Printf(f, goIsFun, goIsFun)
	fmt.Printf(f, maxInt, maxInt)
	fmt.Printf(f, complex, complex)
}
