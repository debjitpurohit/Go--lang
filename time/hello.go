package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study in go")

	// time.Now() returns the current time
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("02-01-2006 Monday 15:04:05")) //// standard format for date and time in go

	///created time for a file
	createTime := time.Date(2020, time.January, 10, 3, 30, 43, 0, time.UTC)
	fmt.Println(createTime)
	fmt.Println(createTime.Format("02-01-2006 Monday 15:04:05 "))

	//to build a app showing time ....{{{     go to git bash    }}}
	// have to create a folder with in create a file with name main.go then go mod init mytime
	//GOOS=windows GOARCH=amd64 go build -o hello-win.exe hello.go ---- 64 bit windows
	//https://freshman.tech/snippets/go/cross-compile-go-programs/
	//https://www.youtube.com/watch?v=HHNVnW3OAQs

}
