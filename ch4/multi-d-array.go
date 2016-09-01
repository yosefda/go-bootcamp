package main

import "fmt"

func main() {
	var a [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			a[i][j] = fmt.Sprintf("row %d - column %d", i, j)
		}
	}

	fmt.Printf("%q\n", a)
	fmt.Println(a[1][1])

	// this will cause compiler error - index out of bound
	//fmt.Println(a[5][5])
}
