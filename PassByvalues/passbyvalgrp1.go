package main

import (
	"fmt"
)

func updateName(x string) {
	x = "Wedge"
}
func mainx() {
	name := "tifa"
	updateName(name)
	fmt.Println(name)
}

//we get tifa and not wedge so when we pass name
//it essentially means its not the defined var its the copy
// of it and not varible itself
// so unless u return and name=: updateName(name) it wont update
