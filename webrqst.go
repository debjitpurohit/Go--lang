package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://lco.dev"

func main() {

	fmt.Println("Web request in golang")
	///GET request
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response type: %T\n", response)
	defer response.Body.Close()                     // caller's responsibility to close the connetion
	databytes, err := ioutil.ReadAll(response.Body) // read the response body

	if err != nil {
		panic(err)
	}
	fmt.Println(string(databytes))

}
