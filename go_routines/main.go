package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter("Hello") // initialize a goroutine
	greeter("World")
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
