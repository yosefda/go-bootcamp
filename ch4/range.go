package main

import "fmt"

func main() {
	cities := []string{"Jakarta", "Bogor", "Sydney", "Adelaide"}
	for idx, val := range cities {
		fmt.Printf("idx: %d \t val: %s\n", idx, val)
	}

	populations := map[string]int{
		"New York":    8336990,
		"Los Angeles": 3900999,
		"Chicago":     1234567,
	}
	for key, val := range populations {
		fmt.Printf("%s has populations of %d \n", key, val)
	}
}
