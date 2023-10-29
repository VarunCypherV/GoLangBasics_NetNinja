package main

import (
	"fmt"
)

func main4() {
	x := 0
	fmt.Println(x)
	for x < 5 {
		fmt.Print(x)
		x++
	}
	fmt.Println()
	for i := 0; i < 8; i++ {
		fmt.Print(i)
	}
	fmt.Println()
	names := []string{"yo", "oy", "yoyo"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	fmt.Println()
	//for in loop
	for index, value := range names {
		fmt.Printf("Pos %v Value %v \n", index, value)
	}
	// if i dont want index or value
	for _, value := range names {
		fmt.Printf("Value %v \n", value)
	}

	for index, value := range names {
		fmt.Printf("Pos %v Value %v \n", index, value)
		value = "newstring"
	}
	fmt.Println(names) //value doesnt update
}
