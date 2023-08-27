package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user-input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating: ")

	rating, _ := reader.ReadString('\n')
	fmt.Println("You entered: " + rating)
}
