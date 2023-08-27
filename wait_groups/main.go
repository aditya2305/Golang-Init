package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals []string

var wg sync.WaitGroup // will be updated to pointers_later
var mut sync.Mutex    // mutex in data [will be updated to pointers]

func main() {
	websiteList := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
		"https://www.youtube.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web) // adding the routines to the waitgroup so that they can be waited on for completion and main is not exited before all routines are completed
		wg.Add(1)             // adding the go_routines count (can be > 1)
	}

	wg.Wait() // usually is present at the end of the method and will wait for all the routines to complete
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	defer wg.Done() // will be sending the signal after all the other tasks are completed

	result, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Oops!")
	}
	mut.Lock() // locking the other writer
	signals = append(signals, endpoint)
	mut.Unlock() // unlocking the other_writer
	fmt.Printf("%d status code endpoint for %s \n", result.StatusCode, endpoint)
}
