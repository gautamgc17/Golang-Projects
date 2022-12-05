package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	// m1 declare arrays
	var fruitList [5]string
	fruitList[0] = "apple"
	fruitList[1] = "mango"
	fruitList[2] = "guava"
	fruitList[4] = "orange"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))               // len is still 5 though we have 4 elements
	fmt.Printf("Type is: %T \n" , fruitList)

	// m2
	var itemList [2]int = [2]int{1}          // no default value, 1st value is 1 and 2nd value initialized with 0
	fmt.Println(itemList)

	// m3 decalare arrays and define array elements
	var vegList = [5]string{"ab", "cd", "ef", "gh", "xyz"}
	fmt.Println(vegList)

	// m4
	newList := [4]string{"a", "b", "c", "d"}
	fmt.Println(newList)

	// m5
	a := [...]int{12, 78, 50}        // ... makes the compiler determine the length
    fmt.Println(a)

	// Arrays in Go are value types and not reference types. 
	// This means that when they are assigned to a new variable, a copy of the original array is assigned to the new variable. 
	// If changes are made to the new variable, it will not be reflected in the original array.
	newa := a
	newa[0] = 99
	fmt.Println(newa)
	fmt.Println(a)
}
