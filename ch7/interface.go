package main

import (
	"fmt"
)

// User User type
type User struct {
	FirstName, LastName string
}

// Name Get user's name
func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

// Namer Everything that has name
type Namer interface {
	Name() string
}

// Greet Greet type
func Greet(n Namer) string {
	return fmt.Sprintf("Hi %s!", n.Name())
}

// Customer Customer type
type Customer struct {
	CustomerID int
	FullName   string
}

// Name Get customer's name
func (c *Customer) Name() string {
	return fmt.Sprint(c.FullName)
}

func main() {
	u := &User{"Matt", "Damon"}
	fmt.Println(Greet(u))
	c := &Customer{1234, "Gina Smith"}
	fmt.Println(Greet(c))
}
