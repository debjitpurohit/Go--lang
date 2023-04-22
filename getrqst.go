package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	performGetRequest()

}

func performGetRequest() {
	const myurl = "http://localhost:5000/get"
	resp, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)
	fmt.Println("content length is :", resp.ContentLength)

	var reponseString strings.Builder
	databytes, _ := ioutil.ReadAll(resp.Body) // read the response body
	byteCount, _ := reponseString.Write(databytes)

	///////////////////// either use builder method or this string conversion method///////////////////

	//builder method
	fmt.Println("Byte count:", byteCount)
	fmt.Println(reponseString)
	fmt.Println(reponseString.String())

	// string conversion method(begineer friendly)
	// fmt.Println(databytes)
	// fmt.Println(string(databytes))

}
