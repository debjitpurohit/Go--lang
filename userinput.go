package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello, World! learnCodeonline.in"
	fmt.Println(welcome) // go is case sensitive

	reader := bufio.NewReader(os.Stdin) // for user input use bufio and os package
	fmt.Println("Enter text: ")

	//comma ok || error ok syntax || try catch method || comma error method
	//just like try catch method
	//input, err := reader.ReadString('\n')
	input, _ := reader.ReadString('\n') // _ is used to ignore the error bcz we have no error here

	fmt.Println("You entered: ", input)
	fmt.Printf("Type of input is %T", input)

}
