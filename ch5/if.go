package main

import "fmt"

func getAnswer() int {
	return 42
}

func main() {
	answer := getAnswer()
	if answer == 42 {
		fmt.Println("Answer is ", answer)
	}

	// if with short statement
	if answer2 := getAnswer(); answer2 == 42 {
		fmt.Println("Answer is ", answer)
	}
}
