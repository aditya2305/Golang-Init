package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang!")

	// data can't be sent to channel unless there is a go_routine
	wg := *&sync.WaitGroup{}
	channel := make(chan int, 1) // creating a channel that handles integer (having 1 as buffer value so that it can listen to multiple values instead of flushing)

	wg.Add(2)

	// receive only go_routines for channel as <-chan
	go func(ch <-chan int, wg *sync.WaitGroup) {
		value, isChannelOpen := <-ch
		fmt.Println("Receiving data from channel: ", value)
		fmt.Println("Is channel open: ", isChannelOpen)
		wg.Done()
	}(channel, &wg)

	// send only go_routines for channel as chan<-
	go func(ch chan<- int, wg *sync.WaitGroup) {
		fmt.Println("Sending data to the channel")
		ch <- 10
		ch <- 15
		close(channel) // closing a channel
		wg.Done()
	}(channel, &wg)

	wg.Wait()
}
