package main

import "fmt"

// User User data structure
type User struct {
	ID             int
	Name, Location string
}

// Greetings Introduce self
func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

// Player Player data structure
type Player struct {
	*User
	GameID int
}

// NewPlayer Create new player
func NewPlayer(id int, name string, location string, gameID int) *Player {
	return &Player{
		&User{
			id,
			name,
			location,
		},
		gameID,
	}
}

func main() {
	player := NewPlayer(1234, "Michael", "Sydney", 67890)
	fmt.Printf("%+v gameID: %v\n", *player.User, player.GameID)
	fmt.Println(player.Greetings())

	// this cause runtime error
	// player2 := &Player{}
	// fmt.Println(player2.Greetings())
}
