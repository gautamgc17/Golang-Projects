package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	contents := "This is initial content inside the file."
	
	file, err := os.Create("./myLogFile.txt")
	checkNilError(err)

	defer file.Close()

	// WriteString writes the contents of the string s to w, which accepts a slice of bytes.
	length, err := io.WriteString(file , contents)
	checkNilError(err)

	fmt.Println("Bytes Written: ", length)
	fmt.Printf("Type of File is: %v \n", file)

	readFile("./myLogFile.txt")
}

func readFile(filepath string){
	// return the contents as bytes
	bytesdata, err := ioutil.ReadFile(filepath)
	checkNilError(err)

	fmt.Println("Bytes form data of file is: ", bytesdata)
	fmt.Println("String form data of file is: ", string(bytesdata))
}

func checkNilError(err error){
	if err != nil{
		// stops the normal execution of the current function
		panic(err)
	}
}
