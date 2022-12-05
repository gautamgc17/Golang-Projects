package main

import (
	"fmt"
	"os"
)

func main() {
	str := "Hello World!!"
	fmt.Println("Hello World" , "Another String")
	fmt.Printf("%s this is string \n" , str )

	fmt.Fprintln(os.Stdout , "Hello World in Golang")
	n , err := fmt.Fprintf(os.Stdout, "%s %s \n" , str, str)
	// 29, nil

	fmt.Print(n , err)
}
