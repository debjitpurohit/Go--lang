package main

import (
	"encoding/json"
	"fmt"
)

// declare a struct
type course struct {
	Name     string `json:"coursename"` // json tag
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // this field will not be exported
	Tags     []string `json:"tags,omitempty"` // if tags is nil then it will not be exported
}

func main() {
	fmt.Println("json handaling")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	// json encode
	lcoCourses := []course{
		{"Go", 100, "Udemy", "1234", []string{"Go", "backend"}},
		{"React", 400, "Udemy", "1234", []string{"React", "frontend"}},
		{"Node", 700, "Udemy", "1234", []string{"Node", "backend"}},
		{"Mongo", 800, "Udemy", "1234", []string{"Mongo", "database"}},
		{"Angular", 500, "Udemy", "1234", nil},
	}
	//package this data as json data
	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") // better json formating better to read
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", finalJson)

}
func DecodeJson() {
	// json decode
	jsonData := []byte(`
		{
			"coursename": "Go",
			"Price": 100,
			"website": "Udemy",
			"tags": [
				"Go",
				"backend"
			]
		}
	`)
	var decoded course
	checkValid := json.Valid(jsonData)
	if checkValid {
		fmt.Println("json is valid")
		json.Unmarshal(jsonData, &decoded)
		fmt.Printf("%#v\n", decoded)

	} else {
		fmt.Println("json is not valid")
	}
	// some cases where you just want to add data to key value

	var decodedMap map[string]interface{}
	json.Unmarshal(jsonData, &decodedMap)
	fmt.Printf("%#v\n", decodedMap)

	for k, v := range decodedMap {
		fmt.Printf("key is %v and value is %v and type is %T\n", k, v, v)
	}

}
