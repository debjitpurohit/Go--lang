package main

import (
	"fmt"
	"sort"
)

func main() {
	//decleare a slice
	fruitList := []string{"Apple", "Orange", "Banana", "Peach"}
	fmt.Printf("type of fruitList is %T \n", fruitList)
	fmt.Println("Fruit List:", fruitList)
	fmt.Println("Fruit List:", len(fruitList))
	//printlist using for loop
	for i := 0; i < len(fruitList); i++ {
		println(fruitList[i])
	}
	//append to a slice
	fruitList = append(fruitList, "Mango", "Grapes")
	fmt.Println("Fruit List:", fruitList)
	fmt.Println("Fruit List:", len(fruitList))
	//remove from a slice
	fruitList = append(fruitList[1:3]) // starts from 1 insted of 0 and ends on (3-1)th index
	//fruitList = append(fruitList[1:])
	//fruitList = append(fruitList[:3])
	fmt.Println("Fruit List:", fruitList)
	fmt.Println("Fruit List:", len(fruitList))

	//////////////use of make and new
	highScore := make([]int, 4)
	highScore[0] = 100
	highScore[1] = 200
	highScore[2] = 300
	highScore[3] = 400
	//highScore[4] = 500 //error

	highScore = append(highScore, 500, 1999, 344)

	fmt.Println("High Score:", highScore)
	fmt.Println(sort.IntsAreSorted(highScore))
	// /////////sorting can be done in slices not in arrays
	sort.Ints(highScore)
	fmt.Println("High Score:", highScore)

	//////use of new
	//new returns a pointer to the memory location
	//make returns a slice
	//newHighScore := new([]int)
	//newHighScore[0] = 100 //error
	//newHighScore[1] = 200 //error
	//newHighScore[2] = 300 //error
	//newHighScore[3] = 400 //error

	//how to remove a value from a slice based on index
	var courses = []string{"python", "java", "c++", "c#", "javascript"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) // tar mane 0 1 print hbe then 2+1 theke start hbe and so on
	fmt.Println(courses)
}
