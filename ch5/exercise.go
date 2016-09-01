package main

import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	// a: 1 coin
	// e: 1 coin
	// i: 2 coins
	// o: 3 coins
	// u: 4 coins
	coinsForUser := func(name string) int {
		userCoin := 0
		for _, char := range strings.ToLower(name) {
			switch char {
			case 'a', 'e':
				userCoin++
			case 'i':
				userCoin += 2
			case 'o':
				userCoin += 3
			case 'u':
				userCoin += 4
			}
		}

		return userCoin
	}

	for _, user := range users {
		userCoin := coinsForUser(user)

		if userCoin > 10 {
			userCoin = 10
		}

		distribution[user] = userCoin
		coins -= userCoin
	}

	fmt.Println(distribution)
	fmt.Println("Coins left: ", coins)
}
