package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Statement")

	// generate random number b/w 0 to n
	t := time.Now().UnixNano()
	fmt.Println(t)
	rand.Seed(t)
	diceNumber := rand.Intn(6) + 1

	// if a matching expression is found, then it exits
	// but if fallthrough is written it moves to next switch case
	switch diceNumber {
	case 1:
		fmt.Println("The dice number is 1")
	case 2:
		fmt.Println("The dice number is 2")
	case 3:
		fmt.Println("The dice number is 3")
	case 4:
		// fallthrough next cases and then terminate
		fmt.Println("The dice number is 4")
		fallthrough
	case 5:
		fmt.Println("The dice number is 5")
		fallthrough
	case 6:
		fmt.Println("The dice number is 6")
	default:
		fmt.Println("Done")
	}
}
