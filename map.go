package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["en"] = "English" //key value pair
	languages["bn"] = "Bangla"
	languages["fr"] = "French"

	fmt.Println("Lists of languages:", languages)
	fmt.Println("Total number of languages:", len(languages))
	fmt.Println("en shorts for", languages["en"])
	delete(languages, "en")
	fmt.Println("Lists of languages:", languages)

	//loops in map
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}
	// for _, value := range languages {
	// 	fmt.Printf("For key v, value is %v \n", value)
	// }
}
