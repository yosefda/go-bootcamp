package main

import "fmt"

func main() {
	cities := make([]string, 3)
	fmt.Printf("%T %q\n", cities, cities)

	cities[0] = "Sydney"
	cities[2] = "Jakarta"
	fmt.Printf("%T %q\n", cities, cities)

	// this will cause runtime error
	// cities[3] = "Bogor"

	names := []string{}
	names = append(names, "Michael")
	fmt.Printf("%T %q\n", names, names)

	names = append(names, "Amy", "Barry")
	fmt.Printf("%T %q\n", names, names)

	otherNames := []string{"Joko", "Painem"}
	// ,,, means we want to iterate the values
	names = append(names, otherNames...)
	fmt.Printf("%T %q\n", names, names)

	// we cant array, [2]string is not the same type as []string
	// specialNames := [...]string{"Stephanov", "Rasyid"}
	// names = append(names, specialNames...)
	// fmt.Printf("%T %q\n", names, names)
}
