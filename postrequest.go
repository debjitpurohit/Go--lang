package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	performPostJsonRequest()
	// PerformPostFormRequest()
}

func performPostJsonRequest() {
	const myurl = "http://localhost:5000/post"

	//fake json payload

	requestBody := strings.NewReader(`
	{
		"firstName": "John",
		"lastName": "Smith",
		"age": 25
	}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

	defer response.Body.Close()

}

func PerformPostFormRequest() {
	const myurl = "http://localhost:5000/postform"

	data := url.Values{}
	data.Add("name", "John Smith")
	data.Add("job", "Programmer")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

	defer response.Body.Close()
}
