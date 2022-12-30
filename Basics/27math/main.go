package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Math , Crypto and Random Number in Golang")

	// var number1 int = 2
	// var number2 float64 = 4.5
	// fmt.Println("The sum:", number1 + int(number2))

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println((rand.Intn(5) + 1))

	// randomness is much more accuracte governed by crypto
	
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
