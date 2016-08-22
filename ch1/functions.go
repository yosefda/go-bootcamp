package main

import (
	"fmt"
)

// This is a function that accepts 2 variables of type int and
// returns an int. The type of function parameters are defined
// after the name.
func add(x int, y int) int {
	return x + y
}

// In this function definition the type of parameters are defined
// once. This means that both parameters have the same type.
func multiply(x, y int) int {
	return x * y
}

// This function accept one parameter and return 2 values.
func location(city string) (string, string) {
	var country, continent string

	switch city {
	case "Sydney", "Adelaide":
		country, continent = "Australia", "Australia"
	case "Jakarta", "Denpasar":
		country, continent = "Indonesia", "Asia"
	}

	return country, continent
}

func main() {
	x, y := 2, 3
	fmt.Printf("%d + %d = %d\n", x, y, add(x, y))
	fmt.Printf("%d * %d = %d\n", x, y, multiply(x, y))

	country, continent := location("Sydney")
	fmt.Printf("Sydney is in %s, %s\n", country, continent)
}
