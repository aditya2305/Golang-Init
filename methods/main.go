package main

import "fmt"

func main() {
	charlie := User{"Charlie", "charlie@email.com", true, 18}
	charlie.GetStatus()
	charlie.NewMail()
}

// defining a struct
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Printf("Is the user active? %v\n", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@email.com"
	fmt.Printf("Email of the user is %v\n", u.Email)
}
