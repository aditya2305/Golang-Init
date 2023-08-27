package main

import "fmt"

func main() {
	loginCount := 7
	var result string

	if loginCount < 10 {
		result = "Login is valid!"
	}

	fmt.Println(result)

	// assigning and checking
	if n := 10; n < 10 {
		fmt.Println("n is less than 10")
	} else {
		fmt.Println("n is greater than 10")
	}
}
