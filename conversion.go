package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our pizza app")
	fmt.Println("please rate our pizza from 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("You rated: ", input)

	//conversion
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // you have to trim the space(\n) or else it will give error{{{{{{   4\n    }}}}}}

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("added 1 to your rating : ", numRating+1)
	}

}
