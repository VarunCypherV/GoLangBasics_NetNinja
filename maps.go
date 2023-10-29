package main

import (
	"fmt"
)

func main7() {
	//all keys same type all value same type
	menu := map[string]float64{
		"soup": 4.99,
		"pie":  7.99,
	}

	fmt.Println(menu, menu["soup"])
	//loops
	for k, v := range menu {
		fmt.Printf("Key : %v , Value : %v\n", k, v)
	}
	menu["soup"] = 5.01
	for k, v := range menu {
		fmt.Printf("Key : %v , Value : %v\n", k, v)
	}
}
