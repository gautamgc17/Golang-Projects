package main

import "fmt"

func main() {
	// There exists anonymous functions, lambda functions and IIFE also
	fmt.Println("Functions and Methods in Golang")

	// order of function call matters
	greeterTwo()
	greeter()

	result := adder(5, 4)
	fmt.Println(result)

	out1, out2 := proAdder(1, 2, 3, 4, 5)
	fmt.Println(out2, out1)
}

// function signature - provide type of arguments as well as return type
func adder(val1 int , val2 int) int {
	return val1 + val2
}

// pass variable length arguments
func proAdder(values ...int) (int , string){
	total := 0
	for _, val := range(values){
		total += val
	}
	return total, "This is the Output"
}

func greeter() {
	fmt.Println("Hello World")
}
func greeterTwo() {
	fmt.Println("Hello World Part 2")
}
