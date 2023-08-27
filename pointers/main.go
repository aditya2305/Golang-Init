package main

import "fmt"

func main() {
	myNumber := 18

	var ptr = &myNumber // referencing to the address

	fmt.Println("Value of the pointer is : ", *ptr)
	fmt.Println("Address of the pointer is : ", ptr)

	*ptr = 45 // changing the value of the pointer

	fmt.Println("Value of the pointer is : ", *ptr)
	fmt.Println("Value of myNumber: ", myNumber)
}
