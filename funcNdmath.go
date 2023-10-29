package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreetings(n string) {
	fmt.Printf("Good morning %v \n", n)
}

// passing func in
func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

//return value

func circleArea(r float64) float64 { //specifiy return type

	return math.Pi * r * r

}

// returns 2 string (down)
// func getInitials(n string) (string,string) {

// }
func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}
	return initials[0], initials[1]
}

func main6() {
	sayGreetings("varun")

	cycleNames([]string{"blake", "raven", "heart"}, sayGreetings)

	a1 := circleArea(10.5)

	fmt.Println(a1)

	i, j := getInitials("tifa lockhart")
	fmt.Println(i, j)

}
