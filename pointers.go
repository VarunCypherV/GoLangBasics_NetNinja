package main

import (
	"fmt"
)

func updateName(x *string) {
	*x = "Wedge"
}
func mainx() {
	name := "tifa"
	m := &name //has its own meomry address and is storaing a value as another memory address
	fmt.Println(&name)
	fmt.Println(name)
	fmt.Println(m)
	fmt.Println(*m)
	// so now printing m and &name is same o/p
	// now if i do dereferencing by *m i get the value the memory address stored in m is storing
	updateName(m)
	fmt.Println(name)
	//now it updates to wedge in original unlike passbyvalgrp1.go

}

//we get tifa and not wedge so when we pass name
//it essentially means its not the defined var its the copy
// of it and not varible itself
// so unless u return and name=: updateName(name) it wont update
