package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	// this reader is a reference or a pointer
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Display the user input")

	// comma ok, comma error syntax
	// err == nil if input is received including the delimiter as well
	input, _ := reader.ReadString('\n')
	fmt.Printf("Type is: %T \n", input)

	// conversion of string to number
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	
	if(err != nil){
		fmt.Println(err)
	} else{
		fmt.Println("Value after sum is:" , numRating + 10)
	}
}
