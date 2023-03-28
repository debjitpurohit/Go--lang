package main

import "fmt"

func main() {
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	// fruitList[2] = "Banana"
	fruitList[3] = "peach"
	//print list just using the variable name
	fmt.Println("Fruit List:", fruitList)
	fmt.Println("Fruit List:", len(fruitList))
	//printlist using for loop
	for i := 0; i < len(fruitList); i++ {
		println(fruitList[i])
	}

	var vegList = [3]string{"carrot", "potato", "tomato"}
	fmt.Println("Veg List:", vegList)

}
