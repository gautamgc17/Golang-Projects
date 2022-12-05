package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
) 

const url = "https://lco.dev"

func main() {
	fmt.Println("Web Requests")

	// Get Requests to a URL
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	// reference of the original response recieved
	fmt.Printf("Response is of type: %T , Status is %v and its Body is: %T \n", response, response.StatusCode, response.Body)
	// caller's responsibility to close the connection
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)

}
