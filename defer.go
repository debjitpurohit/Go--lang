package main

import "fmt"

func main() {
	defer fmt.Println("world1") // by using defer it execute at last of the function
	defer fmt.Println("world2")
	defer fmt.Println("world3") // last in first out
	fmt.Println("debuuuu")
	myDefer() // it have no defer keyword

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) // lifo

	}
}
