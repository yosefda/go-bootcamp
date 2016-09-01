package main

import "fmt"

type Coordinate struct {
	Lat, Long float64
}

func main() {
	// create map using make()
	m := make(map[string]Coordinate)
	fmt.Printf("%#v\n", m)

	m["Yahoo7"] = Coordinate{123.01233, -133.232132}
	m["Google"] = Coordinate{105.23134, -132.453454}
	fmt.Printf("%#v\n", m)

	// retrieve an element
	fmt.Printf("%#v\n", m["Yahoo7"])

	// delete an element using delete()
	delete(m, "Yahoo7")
	fmt.Printf("%#v\n", m)

	// test key is present
	elem, ok := m["Google"]
	fmt.Println(elem, ok)

	// elem will be zero value of the element's type
	// ok will be false
	elem2, ok2 := m["Facebook"]
	fmt.Println(elem2, ok2)
}
