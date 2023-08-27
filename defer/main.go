package main

import "fmt"

func main() {

	// defer will follow the LIFO order for execution
	defer fmt.Println("world")
	defer fmt.Println("Another")
	defer fmt.Println("This")
	fmt.Println("Hello")
}
