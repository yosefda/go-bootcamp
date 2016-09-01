package main

import "fmt"

func main() {
	// standard for loop
	sum1 := 0
	for i := 1; i <= 10; i++ {
		sum1 += i
	}
	fmt.Println("sum1 is ", sum1)

	// without pre and post conditions
	sum2 := 1
	for sum2 <= 50 {
		sum2 += sum2
	}
	fmt.Println("sum2 is ", sum2)

	// as a while loop
	sum3 := 1
	for sum3 <= 60 {
		sum3 += sum3
	}
	fmt.Println("sum3 is ", sum3)
}
