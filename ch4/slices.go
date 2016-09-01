package main

import "fmt"

func main() {
	s := []int{3, 4, 5, 6, 7}
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	fmt.Println(s[3:])
	fmt.Println(s[2:5])
	fmt.Println(s[:4])
}
