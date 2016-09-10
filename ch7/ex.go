package main

import (
	"fmt"
)

// ErrNegativeSqrt Negative param given for Sqrt
type ErrNegativeSqrt float64

// Error implicitly implements error
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot sqrt negative number: %g", float64(e))
}

// Sqrt Sqrt function
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.
	for i := 0; i < 10; i++ {
		z = z - ((z*z)-x)/(2*z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(5))
	fmt.Println(Sqrt(0))
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(-4))
}
