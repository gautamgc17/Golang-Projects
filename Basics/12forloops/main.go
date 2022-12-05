package main

import "fmt"

func main() {
	fmt.Println("For Loops")

	sample := []string{"ab", "cd", "xy", "ef"}
	for d := 0; d < len(sample); d++ {
		fmt.Println(sample[d])
	}
	fmt.Println()

	// to iterate over elements over data structures
	for i := range sample {
		fmt.Println(sample[i])
	}
	fmt.Println()

	// can provide one or two values
	for idx, key := range sample {
		fmt.Printf("Key %v and Value %v \n", idx, key)
	}
	fmt.Println()

	// incase of strings, 2nd val returned is int32 ASCII eq.
	str := "sample"
	for idx, val := range str {
		fmt.Printf("Key %v and Value %v \n", idx, val)
	}
	fmt.Println()

	// break, continue, goto statement (jump tp label statement)
	i := 1
	for i < 10 {
		if i == 3 {
			// skip a phase and provide that phase
			i++
			continue
		}
		if i == 5 {
			goto label
		}
		fmt.Println(i)
		i++
	}
	
	// define after loop statement
	label:
		fmt.Println("Yo statement jumped here")


	for d := 0; d < len(sample); d++ {
		if(d==2){
			continue
		}
		fmt.Println(sample[d])
	}
}
