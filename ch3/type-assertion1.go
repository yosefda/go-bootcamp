package main

import (
	"fmt"
	"time"
)

// interface{} type is kinda like Object in Java
func timeMap(y interface{}) {
	// type assert y to be map[string]interface
	z, ok := y.(map[string]interface{})

	// if yes then add updated_at key
	if ok {
		z["updated_at"] = time.Now()
	}
}

func main() {
	foo := map[string]interface{}{
		"Matt": 42,
	}

	bar := map[int]string{
		121: "Michael",
	}

	// updated_at will be added
	timeMap(foo)
	fmt.Printf("%T (%v)\n", foo, foo)

	// updated_at wont be added
	timeMap(bar)
	fmt.Printf("%T (%v)\n", bar, bar)
}
