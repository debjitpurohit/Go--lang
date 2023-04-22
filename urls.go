package main

import (
	"fmt"
	"net/url"
)

const myurl string = "http://lco.dev:3000/learn?course=python&lesson=1"

func main() {
	fmt.Println(myurl)
	//parsing

	// parts of a url
	result, _ := url.Parse(myurl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	//query params
	qparams := result.Query()
	fmt.Printf("type of query parameters : %T\n", qparams)
	fmt.Println(qparams["course"])

	for key, value := range qparams {
		fmt.Printf("params key is %s and value is %s\n", key, value)
	}

	/// building a url
	partsOfUrl := &url.URL{
		Scheme:   "http",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "course=go&lesson=3",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
