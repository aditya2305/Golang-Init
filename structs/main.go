package main

import "fmt"

func main() {
	charlie := User{"Charlie", "charlie@email.com", true, 18}
	fmt.Println(charlie.Name)
	fmt.Printf("Details for %v are: %+v", charlie.Name, charlie)
}

// defining a struct
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
