package main

import (
	"fmt"
	"sort"
)

func main() {
	// A slice is a convenient, flexible and powerful wrapper on top of an array. 
	// Slices do not own any data on their own. They are just references to existing arrays.
	fmt.Println("Learning Slices")

	// slices have expandable memory
	var fruitList = []string{}
	fmt.Printf("Type is %T \n" , fruitList)

	// slice = append(slice, elem1, elem2) - appends elements to the end of a slice
	fruitList = append(fruitList, "hello" , "world" , "go" , "lang" , "python")
	fmt.Println(fruitList)
	fruitList = append(fruitList[:2])
	fmt.Println(fruitList)
	// length and capacity of slice
	fmt.Printf("length of slice %d capacity %d \n", len(fruitList), cap(fruitList))

	// slice = append(slice, anotherSlice...)
	fruitList = append(fruitList[:2] , fruitList[2:4]...)
	fmt.Println(fruitList)

	// another syntax - make() allocates and initializes of given type
	highScores := make([]int , 4)
	highScores[0] = 900
	highScores[1] = 165
	highScores[2] = 483
	highScores[3] = 650                     // highScores[4] = 200 will throw error
	// reallocation of memory takes place
	highScores = append(highScores, 199, 200)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "java", "python", "swift"}
	var index int = 2
	courses = append(courses[:index] , courses[index+1:]...)
	fmt.Println(courses)

}
