package main

//main is global so any declared outside funcs in main
//is global to all in directory
import "fmt"

var score = 55

func main() {

	sayHello("mario")
	for _, v := range points {
		fmt.Println(v)
	}
	showScore()
}

//run go run utilimain.go utility.go
