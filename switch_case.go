package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value is one and u can open")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move 3 spot")
		fallthrough // used to print 3 ,4 as well as 5
	case 4:
		fmt.Println("you can move 4 spot")
		fallthrough // print 4 as well as 5
	case 5:
		fmt.Println("you can move 5 spot")
	case 6:
		fmt.Println("you can move the 6 spot and roll the dice again")
	default:
		fmt.Println("WTF!")
	}
}
