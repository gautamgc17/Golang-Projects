package main

import "fmt"

func main() {
	fmt.Println("If Else-If Else")

	if 9%2 == 0{
		fmt.Println("Even")
	} else{
		fmt.Println("Odd")
	}

	// assigning value on the run and checking (useful in web requests)
	if num := 3; num < 2 {
		fmt.Println("Value less than 2")
	} else if (num >= 2 && num < 10) {
		fmt.Println("Between 2 and 10")
	} else{
		fmt.Println("Value above than 10")
	}
}
