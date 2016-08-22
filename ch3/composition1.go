package main

import "fmt"

// User Data structure that defines a user
type User struct {
	ID             int
	Name, Location string
}

// Greetings Get user intro text
func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s!\n", u.Name, u.Location)
}

// Player Data structure that defines a player.
// Player is composed from User.
type Player struct {
	User   // composed from User
	GameID int
}

func main() {
	// create player using dot notation
	player1 := Player{}
	player1.ID = 1
	player1.Name = "Michael"
	player1.Location = "Kingsford"
	player1.GameID = 12345

	fmt.Printf("Player 1: %v\n", player1)
	fmt.Println(player1.Greetings())

	// create player using struct literal
	player2 := Player{
		User: User{ //apparently need to specify User:
			ID:       2,
			Name:     "Nathan",
			Location: "Jati Asih",
		},
		GameID: 6789,
	}

	fmt.Printf("Player 2: %v\n", player2)
	fmt.Println(player2.Greetings())
}
