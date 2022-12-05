package main

import "fmt"


// No Inheritance, No Super keyword or Parent/Child
type User struct {
	// exportable properties since they start with capital letter
	Name   string
	Email  string
	Status bool
	Age    int
}

// Methods - Functions associated with objects
// pass struct object as parameter
func (u User) GetMethod() {
	fmt.Println("User status is: ", u.Status)
}

func (u User) NewMail(){
	u.Email = "abcd@xyz.com"
	fmt.Println("User email is: ", u.Email)
}

func main() {
	fmt.Println("Structs")

	user1 := User{"sample", "sample@co.in" , true , 21}
	fmt.Println(user1)

	// name-values are printed with %+v syntax
	fmt.Printf("Details of user1 are: %+v \n" , user1)
	fmt.Printf("Name of user1 is: %v \n" , user1.Name)

	user1.GetMethod()
	user1.NewMail()

	// no change in original mail, since a copy of object or struct is passes here
	fmt.Printf("Details are: %+v", user1)
}
