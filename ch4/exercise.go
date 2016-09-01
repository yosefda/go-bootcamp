package main

import "fmt"

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}

func main() {
	// find out the longest string
	max_len := 0
	for _, name := range names {
		strlen := len(name)
		if strlen > max_len {
			max_len = strlen
		}
	}

	// iterate names and put into bucket
	output := make([][]string, max_len)
	for _, name := range names {
		idx := len(name) - 1
		output[idx] = append(output[idx], name)
	}

	fmt.Println(output)
}
