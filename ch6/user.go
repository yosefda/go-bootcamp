package main

import (
	"fmt"
)

// User User type
type User struct {
	FirstName, LastName string
}

// Greeting Greet the user
func (u *User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

func main() {
	u := User{"Sherlock", "Holmes"}
	fmt.Println(u.Greeting())
}
