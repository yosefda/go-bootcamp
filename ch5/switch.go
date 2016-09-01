package main

import (
	"fmt"
	"time"
)

func main() {
	// standard switch
	now := time.Now().Unix()
	mins := now % 2
	switch mins {
	case 0:
		fmt.Println("even")
	case 1:
		fmt.Println("odd")
	}

	// using expression in the case statement
	num := 3
	v := num % 2
	switch v {
	case 0:
		fmt.Println("even")
	case 3 - 2: // an expression
		fmt.Println("odd")
	}

	// multiple values in the case statement
	score := 7
	switch score {
	case 0, 1, 3:
		fmt.Println("Terrible")
	case 4, 5:
		fmt.Println("Mediocore")
	case 6, 7:
		fmt.Println("Not bad")
	case 8, 9:
		fmt.Println("Almost perfect")
	case 10:
		fmt.Println("Hmmm did you cheat?")
	default:
		fmt.Println(score, " is off the chart")
	}

	// using fallthrough to execute all statements after
	// a match
	n := 2
	switch n {
	case 0:
		fmt.Println("is zero")
		fallthrough
	case 1:
		fmt.Println("<= 1")
		fallthrough
	case 2:
		fmt.Println("<= 2")
		fallthrough
	case 3:
		fmt.Println("<= 3")
		fallthrough
	default:
		fmt.Println("try again")
	}

	// using break to exit the switch processing
	n2 := 2
	switch n2 {
	case 0:
		fmt.Println("is zero")
		fallthrough
	case 1:
		fmt.Println("<= 1")
		fallthrough
	case 2:
		fmt.Println("<= 2")
		if time.Now().Unix()%2 == 0 {
			fmt.Println("Breaking out!!!")
			break
		}
		fallthrough
	case 3:
		fmt.Println("<= 3")
		fallthrough
	default:
		fmt.Println("try again")
	}
}
