package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gautamgc75/mongoapi/router"
)

func main() {
	fmt.Println("Welcome to MongoDB API")

	r := router.Router()
	fmt.Println("Starting the server.....")
	log.Fatal(http.ListenAndServe(":4000" , r))
	fmt.Println("Listening at the port 4000")
}
