package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Performing different Web Requests")

	// const myurl = "http://localhost:8000/get"
	// PerformGetRequest(myurl)

	// const myurl = "http://localhost:8000/post"
	// PerformPostJsonRequest(myurl)

	const myurl = "http://localhost:8000/postform"
	PerformPostFormRequest(myurl)
}

func PerformGetRequest(myurl string) {
	res, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("Status Code:", res.StatusCode)
	fmt.Println("Content Length:", res.ContentLength)

	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))

	// Build string efficiently using Write method, provide more capabilities and minimize copying
	// Write appends the contents of byte content to b's buffer, so it holds the raw data all the time
	var responseString strings.Builder
	bytecount, _ := responseString.Write(content)
	fmt.Println("Bytes count:", bytecount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest(myurl string) {
	// Fake Json Payload for Post
	// NewReader returns a new Reader reading from s
	requestBody := strings.NewReader(`
		{
			"course":"Golang",
			"price":"699"
		}
	`)
	// Post issues a POST to the specified URL and specified reader body
	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Content:", string(content))
	fmt.Println("Status: ", response.StatusCode)
}

func PerformPostFormRequest(myurl string) {
	// Formdata
	// Values maps a string key to a list of values. It is typically used for query parameters and form values
	data := url.Values{}
	data.Add("firstname", "gautam")
	data.Add("lastname", "devilliers")
	data.Add("role", "wicket-kepper batsman")
	data.Add("email", "gc17@gmail.com")

	resp, err := http.PostForm(myurl , data)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Status Code:" , resp.StatusCode)
	fmt.Println("Content is:" , string(content))

}
