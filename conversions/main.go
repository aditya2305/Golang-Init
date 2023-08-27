package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a rating between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	rating, _ := reader.ReadString('\n')

	fmt.Println("You entered: " + rating)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println("Error: ", err)
	} else if numRating < 1 || numRating > 5 {
		fmt.Println("Error: Rating must be between 1 and 5")
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
