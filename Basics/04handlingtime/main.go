package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Welcome to working with Time and Date")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	createdAt := time.Date(2022, time.November, 23, 11, 28, 10, 45, time.UTC)
	fmt.Println(createdAt)
	fmt.Println(createdAt.Format("02/01/2006 15:04:05 Monday"))

	fmt.Println(runtime.NumCPU())
}
