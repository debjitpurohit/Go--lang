package main

import "fmt"

func main() {
	// days := []string{"sunday", "tuesday", "wednesday"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println("the day is", days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	rougeValue := 1
	for rougeValue < 10 {

		if rougeValue == 2 {
			goto lco
		}

		if rougeValue == 5 {

			rougeValue++
			continue
			// break
		}

		fmt.Println("value is ", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("learncodeonline.in")
}
