package main

import "fmt"

//decleare const
const LoginToken string = "jhhgjhghjgjjf" // global public variable with capital L

func main() {
	// %T for find the type of variable
	var username string = "learnCodeonline.in"
	var age int = 10
	var isCool bool = true
	var smallval uint8 = 255
	var smallfloat float32 = 255.123455789
	fmt.Println(username)                              // go is case sensitive // fp
	fmt.Printf("variable is of type: %T \n", username) ////
	fmt.Println(age)                                   // go is case sensitive
	fmt.Printf("variable is of type: %T \n", age)      /////printf
	fmt.Println(isCool)                                // go is case sensitive /// println
	fmt.Printf("variable is of type: %T \n", isCool)
	fmt.Println(smallval) // go is case sensitive
	fmt.Printf("variable is of type: %T \n", smallval)
	fmt.Println(smallfloat) // go is case sensitive
	fmt.Printf("variable is of type: %T \n", smallfloat)
	////////alias///////////////
	//default values and some alias
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("another variable is of type: %T \n", anothervariable)
	var anotherstring string
	fmt.Println(anotherstring)

	//implicit declaration lexer help in this
	var website = "learncodeonline.in"
	var values = 1
	fmt.Println(website)
	fmt.Println(values)
	fmt.Printf("variable is of type: %T \n", website)
	fmt.Printf("variable is of type: %T \n", values)

	//no var keyword
	//short hand declaration
	// := is used for declaration and assignment
	nuberofuser := 100
	fmt.Println(nuberofuser)
	fmt.Printf("variable is of type: %T \n", nuberofuser)
	numberoffloat := 100.123456789
	fmt.Println(numberoffloat)
	fmt.Printf("variable is of type: %T \n", numberoffloat)
	fmt.Println(LoginToken)

}

/*
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers (5 values after decimal)
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
*/
