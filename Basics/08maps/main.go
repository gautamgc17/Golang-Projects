package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// var mymap map[string]interface{}
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["abc"] = "sample"
	languages["eg"] = "example"
	fmt.Println(languages)
	fmt.Println(languages["JS"])

	// The delete built-in function deletes the element with the specified key
	delete(languages , "eg")

	// loop through the elements - %T - for type and %v - for value
	for _ , val := range(languages){
		fmt.Printf("For Key , Value is %v\n", val)
	}
}
