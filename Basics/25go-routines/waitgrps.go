// Waitgroups (come from sync package)
// For completion of execution of goroutines

// Mutex
// Lock the memory until write operation is done

package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}

var waitGrp sync.WaitGroup   // generally pointer
var mut sync.Mutex           // generally pointer


func main() {
	fmt.Println("WaitGroups")

	start := time.Now()

	websites := []string{
		"https://lco.dev",
		"https://go.dev",
 		"https://google.com",
 		"https://fb.com",
 		"https://github.com",
	}

	for _, web := range websites {
		go getStatusCode(web)
		waitGrp.Add(1)
	}

	// Wait blocks until the WaitGroup counter is zero, stop main func from exiting till all goroutines are done
	// Usually at the end of main function
	waitGrp.Wait()

	fmt.Println("Time taken: " , time.Since(start))
}


func getStatusCode(endpoint string) {

	// Done decrements the WaitGroup counter by one
	defer waitGrp.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("There was some error")
	} else {
		
		// put a lock over memory when multiple goroutines might try to access same code section
		mut.Lock()
		signals = append(signals, endpoint)
		// unlock
		mut.Unlock()
		
		fmt.Printf("%d status code for %s \n", res.StatusCode, endpoint)
	}
}
