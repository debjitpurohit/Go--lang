package main

import "fmt"

func greeter() {
	fmt.Println("namaste from debjit purohit")
}

func adder(n int, m int) int {
	return n + m

}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val

	}
	return total
}

func addermsg(n int, m int) (int, string) {
	return n + m, "this is a return message "

}

func main() {
	fmt.Println("welcome to go lang")
	greeter()

	result := adder(2, 3)
	fmt.Println("result is", result)

	proresult := proAdder(2, 3, 4, 5, 6, 7, 7, 8, 89, 9, 0, 6, 5, 6, 7, 8, 9, 9)
	fmt.Println("value of proresult is", proresult)

	adder2nd, adder2ndmsg := addermsg(4, 8)
	fmt.Println(adder2nd)
	fmt.Println(adder2ndmsg)

	// func greetertwo(){
	// 	fmt.Println("another method") /////////not allow
	// }
}
