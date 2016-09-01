package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)
	fmt.Printf("%q\n", a)
	fmt.Printf("%s\n", a)

	// set the array entries as we declare the array
	b := [2]string{"hello", "again"}
	fmt.Println(b)

	// not specifying the length
	c := [...]string{"aye", "sir!!"}
	fmt.Println(c)
}
