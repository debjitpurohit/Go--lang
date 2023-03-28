package main

import "fmt"

func main() {
	// var one int = 1

	//var myPointer *int // *int is a pointer to an int
	//println(myPointer) // 0x0 as it is not initialized<nil>

	myNumber := 23
	var ptr = &myNumber // & is the address of operator

	fmt.Println("adress of mynumber", ptr)
	fmt.Println("value of mynumber", *ptr)

	*ptr = *ptr * 2
	fmt.Println("adress of mynumber after multification", ptr)
	fmt.Println("value of mynumber after multification", *ptr)
	fmt.Println("adress of mynumber after multification", myNumber)

}
