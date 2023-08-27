package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Test for race_conditions")

	wg := &sync.WaitGroup{} // storing a wait_group pointer
	mut := &sync.RWMutex{}

	var data = []int{0}

	wg.Add(3) // total_number of go-routines
	// IIFE in golang
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.Lock()
		data = append(data, 1)
		mut.Unlock()
		fmt.Println("One!")
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.Lock()
		data = append(data, 2)
		mut.Unlock()
		fmt.Println("Two!")
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.Lock()
		data = append(data, 3)
		mut.Unlock()
		fmt.Println("Three!")
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		mut.RLock() // adding mutex locking for reading
		fmt.Println("Reading: ", data)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(data)
}
