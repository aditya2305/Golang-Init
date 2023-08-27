package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Methods!")

	presentTime := time.Now()

	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 "))

	createdDate := time.Date(2017, time.August, 15, 17, 18, 14, 0, time.UTC)
	fmt.Println(createdDate)
}
