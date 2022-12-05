// Goroutines – Concurrency in Golang - Do not communicate by sharing memory, instead share memory by communicating.
// A Goroutine is a function or method which executes independently and simultaneously in connection with any other Goroutines present in your program.
// All the Goroutines are working under the main Goroutines if the main Goroutine terminated, then all the goroutine present in the program also terminated.
// Goroutine always works in the background.
// The control does not wait for Goroutine to complete their execution just like normal function they always move forward to the next line after the Goroutine call 
// and ignores the value returned by the Goroutine.

// How are Goroutines implemented?
// Goroutines are lightweight, costing little more than the allocation of stack space. The stacks start small and grow by allocating and freeing heap storage as req.
// Internally goroutines act like coroutines that are multiplexed among multiple operating system threads.

// Advantages of Goroutines
// Goroutines are cheaper than threads.
// Goroutine are stored in the stack and the size of the stack can grow and shrink according to the requirement of the program. But in threads, the size of the stack is fixed.
// Goroutines can communicate using the channel and these channels are specially designed to prevent race conditions when accessing shared memory using Goroutines.
// Suppose a program has one thread, and that thread has many Goroutines associated with it. If any of Goroutine blocks the thread due to resource requirement then all the remaining Goroutines will assign to a newly created OS thread. All these details are hidden from the programmers.

// Mutex
// Lock the memory until write operation is done

// Race condition
// Multiple goroutines trying to write into same memory space at the same time

package main

import (
	"fmt"
	"time"
)

func greeter(s string) {
	for i := 0; i < 5; i++ {
		// Sleep pauses the current goroutine for at least the duration d ; main code gets sleep for 2 millisecond
		fmt.Println(s)
		time.Sleep(2*time.Millisecond)
	}
}

func main() {
	// fire up the thread to execute this function with Hello keyword but we never waited for that thread to come back
	go greeter("Hello")
	greeter("World")
}
