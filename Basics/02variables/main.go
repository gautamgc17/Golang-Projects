package main

import "fmt"

// constants, public methods or variables start with capital letter
const LoginToken string = "qwertyuiop"
const Century = 99;

func main() {
	fmt.Println(LoginToken , "\n" , Century)


	// declare variable
	var username string = "Raman"
	// print according to default formatting
	fmt.Println("Variables, Types and Constants")
	// print according to specified formatting
	fmt.Printf("Variable is of Type: %T \n", username)


	var isLoggedIn bool = false
	fmt.Printf("Variable is of Type: %T \n", isLoggedIn)
	var smallVal float32 = 25.123456789
	fmt.Printf("Variable is of Type: %T \n", smallVal)
	var smallAns uint8 = 255
	fmt.Printf("Variable is of Type: %T \n", smallAns)


	// no garbage value is set, 0 for int/float and empty string for strings
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Type is: %T \n" , anotherVar)
	var checkVar bool
	fmt.Println(checkVar)


	// implicit type, cannot re-assign new type
	var newVar = "implicit"
	newVar = "new-example"
	fmt.Println(newVar)
	fmt.Printf("Type is: %T \n", newVar)


	// no var style, use of walrus operator, can be used inside methods only not in global scope
	variable := 85
	fmt.Printf("Type is: %T \n", variable)


	// mutiple variable declaration
	var n1, n2, n3 int = 10, 20, 30
	var name, age, fees = "gautam", 21, 65.2500
	fmt.Println(n1 , n2 , n3)
	fmt.Printf("Type is: %T , %T , %T" , name , age , fees)
}
