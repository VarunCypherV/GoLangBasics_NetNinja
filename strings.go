package main

//importing multiple
import (
	"fmt"
	"strings" //for string methods
)

func main2() {

	greeting := "heyyy yooooo"
	fmt.Println(strings.Contains(greeting, "heyyy"))
	fmt.Println(strings.ReplaceAll(greeting, "heyyy", "hi"))
	//they return new string and not alter original
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "yooo"))
	fmt.Println(strings.Split(greeting, " "))
}
