// Deffered functions are invoked immediately before the surrounding function returns
// In the reverse order they were deferred (LAST IN FIRST OUT)

package main

import "fmt"

func main() {
	fmt.Println("Defer Statement")

	// LIFO
	defer fmt.Println("Defer Statements")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello World")
	myDefer()
}

func myDefer(){
	for i:=0; i<5; i++{
		defer fmt.Println(i)
	}
}
