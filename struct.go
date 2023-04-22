package main

import "fmt"

func main() {
	fmt.Println("struct in go lang")
	/// no inheritance in go lang
	// capatilize the first letter of the struct name
	//declare a struct
	debu := person{"Debjit purohit", 30, "debjit.shilpi@gmail.com"}
	fmt.Println(debu)
	fmt.Printf("Debu's details are\n%+v\n", debu) // details with parameter
	fmt.Printf("Debu's name  is %v\n", debu.Name)
	fmt.Printf("Debu's age is %v\n", debu.Age)
	fmt.Printf("Debu's email  is %v\n", debu.Email)

}

//declare a struct class
type person struct {
	Name  string
	Age   int
	Email string
}
