package main

import "fmt"

func main() {

	strings := []string{"a", "b", "c", "d", "e"}
	fmt.Println(strings)

	// for d := 0; d < len(strings); d++ {
	// 	fmt.Println(strings[d])
	// }

	for _, value := range strings {
		fmt.Println(value)
	}

	rogueValue := 1
	for rogueValue < 10 {
		if rogueValue == 4 {
			goto loco
		}
		fmt.Println(rogueValue)
		rogueValue++
	}

loco:
	fmt.Println("This is loco!")
}
