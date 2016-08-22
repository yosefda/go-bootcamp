package main

import (
	"fmt"
)

func main() {
	// Note that in go variable names can't contain _.
	name, age, location := "Michael", 1, "Sydney"
	fmt.Printf("Name %s age is %d year(s) and lives in %s\n", name, age, location)

	// Variable contains a function.
	sayHello := func(name string) {
		fmt.Printf("Hello %s!\n", name)
	}

	sayHello(name)
}
