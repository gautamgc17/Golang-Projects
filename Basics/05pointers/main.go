package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	// declare a pointer and default value is nil
	var ptr *int
	fmt.Println(ptr)

	//  reference variable holding address of variable
	myNum := 20
	var ptr2 = &myNum
	fmt.Println(ptr2)
	fmt.Println(*ptr2)

	*ptr2 += 5
	fmt.Println(ptr2)
	fmt.Println(*ptr2)
}