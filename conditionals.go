package main

import (
	"fmt"
)

func main5() {
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("in first")
	} else if age < 40 {
		fmt.Println("in second")
	} else {
		fmt.Println("in last")
	}
	names := []string{"mario", "luigi", "doom", "boom"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index == 2 {
			fmt.Println("break at pos", index)
			break
		}
		fmt.Printf("value at pos %v is %v\n", index, value)
	}
}
