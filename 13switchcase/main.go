package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Welcome to switch in Go!")

	rand.Seed(time.Now().UnixNano())
	diceValue := rand.Intn(6) + 1

	fmt.Println("Dice Value:", diceValue)

	switch diceValue {
	case 1:
		fmt.Println("You got one")
	case 2:
		fmt.Println("You got two")
	case 3:
		fmt.Println("You got three")
	case 4:
		fmt.Println("You got four")
	case 5:
		fmt.Println("You got five")
	case 6:
		fmt.Println("You got six")
	default:
		fmt.Println("You got something else")
	}
}
