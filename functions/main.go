package main

import "fmt"

func greeter() {
	fmt.Println("Hello, World!")
	res := adder(4, 5)
	proRes, message := proAdder(1, 2, 3, 4, 5)

	fmt.Println(res)
	fmt.Println(message, proRes)
}

func adder(first int, second int) int {
	return first + second
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}

	return total, "Pro res value:"
}

func main() {
	greeter()
}
