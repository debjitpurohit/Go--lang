package main

func main() {
	loginCount := 2
	var result string
	if loginCount > 10 {
		result = "You are an expert user."
	} else if loginCount > 5 {
		result = "You are an intermediate user."
	} else {
		result = "You are a beginner."

	}
	println(result)
	///assinging and checking at same time with in if else
	if num := 10; num > 0 {
		println("Number is positive")
	} else {
		println("Number is negative")
	}

	// if err!= nil{
	// 	//do something
	// }else {
	// 	//do something
	// }

}
