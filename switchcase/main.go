package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1 // will generate a number between 1 and 6

	switch diceNumber {
	case 1:
		fmt.Println("You got a one")
	case 2:
		fmt.Println("You got a two")
	case 3:
		fmt.Println("You got a three")
	case 4:
		fmt.Println("You got a four")
	case 5:
		fmt.Println("You got a five")
	case 6:
		fmt.Println("You got a six")
	default:
		fmt.Println("You got an invalid number")
	}
}
