// CHANNELS and DEADLOCKS - Pipeline through which multiple goroutines talk to each other, not knowing what happening inside

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels and Deadlocks")

	wg := &sync.WaitGroup{}

	// Unbuffered channel of integer type
	myCh := make(chan int , 10)

// 	Unbuffered := make(chan int)    // Unbuffered channel of integer type
//  buffered := make(chan int, 10)	// Buffered channel of integer type


	// fatal error: all goroutines are asleep - deadlock!
	// Passing value is only allowed by channel if someone is listening to it
	// myCh <- 5              // set value in channel
	// fmt.Println(<-myCh)    // get value

	wg.Add(2)
	// READ ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// two listeners/producers for consuming two values or use buffer channels
		// fmt.Println(<- myCh)
		// fmt.Println(<- myCh)

		val, isChanelOpen := <- myCh
		if isChanelOpen{
			fmt.Println(val)
		} else{
			fmt.Println("Listening to Closed Channel")
		}
		
		wg.Done()
	}(myCh , wg)


	// SEND ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// myCh <- 5
		// myCh <- 6

		// The close built-in function closes a channel, which must be either bidirectional or send-only. 
		// It should be executed only by the sender, never the receiver.
		close(myCh)
		// myCh <- 0     // writing to closed channel not allowed

		wg.Done()
	}(myCh , wg)

	wg.Wait()
}
