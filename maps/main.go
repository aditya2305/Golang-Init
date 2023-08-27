package main

import "fmt"

func main() {
	languages := make(map[string]string) // [key]value

	languages["Go"] = "Golang"
	languages["Python"] = "Python"
	languages["Ruby"] = "Ruby"
	languages["Java"] = "Java"

	delete(languages, "Java")
	// looping a map
	for key, value := range languages {
		fmt.Printf("For the key %v, value is: %v\n", key, value)
	}
}
