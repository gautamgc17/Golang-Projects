package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghjbj4569abc"

func main() {
	fmt.Println("Welcome to handling URLs")
	
	// Parse parses a raw url into a URL structure
	result, _ := url.Parse(myurl)

	fmt.Printf("Type is: %T \n", result)     // *url.Urls
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	// Key-Value of query parameters
	qparams := result.Query()
	fmt.Printf("The type of query params is: %T \n", qparams)     // *url.Values
	for key, val := range qparams{
		fmt.Println(key , "\t" , val[0])
	}

	// Constructing URL
	partsOfUrl := *&url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawQuery: "user=hitesh",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}