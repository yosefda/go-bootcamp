package main

import (
	"fmt"
	"math"
)

// MyFloat Type alias for float64
type MyFloat float64

// Abs  Return absolute value
func (mf MyFloat) Abs() float64 {
	f := float64(mf)
	if f < 0 {
		return -f
	}

	return f
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
