// Race condition
// Multiple goroutines trying to write into same memory space at the same time

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Conditions in Golang")

	score := []int{0}
	wg := &sync.WaitGroup{}    // pointer
	mut := &sync.Mutex{}


	wg.Add(3)
	go func (wg *sync.WaitGroup , mut *sync.Mutex) {
		fmt.Println("One R")
		mut.Lock()
		// Have RLock() while accessing that resource
		score = append(score , 1)
		mut.Unlock()
		// Schedule the call to Done to tell main we are done
		wg.Done()
	} (wg , mut)


	go func (wg *sync.WaitGroup , mut *sync.Mutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score , 2)
		mut.Unlock()
		wg.Done()
	} (wg , mut)

	go func (wg *sync.WaitGroup , mut *sync.Mutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score , 3)   
		mut.Unlock()
		wg.Done()
	} (wg , mut)


	wg.Wait()
	fmt.Println(score)
}
