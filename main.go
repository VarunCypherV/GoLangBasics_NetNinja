// Compiled language so faster than python
// static typed : var types checked by compiler
// strongly typed : var types cant ccchange after decclared
// unlike typed
// Built in test lib like react
// Object oriented alng i own way but not classes
// but uses custom types structsand interfaces
//Go is a pass by value language , meaning when passing varibale as arg in fucn , go makes
//copies of values for the func

// types of var : str , int , float , bool ,array , structs grp 1
// slcies maps and funcs grp 2

package main

// every go file is satrted with package
// if main then go will create a .exe file on executing

import "fmt" //for printing : formating string and printing in console
//functions begin with func
// " func main " is the entry point of app SAME NAME and only 1 in whole app
func main1() {
	//functions like println with begin captical
	// cuz while exporting in package u need capital letter
	fmt.Println("Hello, world!")

	//STRINGS AND VARIABLE DECLARING
	var nameOne string = "Hello, world!"
	//no single quotes
	//in go if u declare u have to use else error even for pacckage or funcs
	var nameTwo = "Hello 2" //infers type automatically

	var nameThree string
	nameThree = "Hello, three "

	//def val for string if not declared is empty

	namefour := "lesgo"
	//short order first time intialization varible
	//infers variable and type but cant use this shorthand method outside func
	fmt.Println(nameOne, nameTwo, nameThree, namefour)

	//INTEGERS
	var age int = 20
	var age2 = 2
	age3 := 44

	//Bits and Memory
	var bitno int8 = 25 //int16 etc
	//integer that can be represented in 8 bits can be used here
	var unsignedint uint8 = 25
	//cant have negative no and can be rep in 8 bit , now this range is more than int8 cuz -ve no not included
	fmt.Println(age, age2, age3, bitno, unsignedint)

	//FLOAT : telling bitsize is compulsory unlike ints
	var score float32 = 25.63
	var scoretwo float64 = 342424.35345353
	scorethree := 1.5 //it is taken as float64 as its the defualt
	fmt.Println(score, scoretwo, scorethree)

	//PRINTING FUNCS
	fmt.Print("hello, ") //no endl here while println moves to new line
	fmt.Print("heu \n")
	fmt.Print("my age is ", age, " and my name is", nameOne, "\n")

	fmt.Printf("my age is %v and my name is %v \n", age, nameOne)
	//format specifier default for variable
	//is is of the form %_
	fmt.Printf("my age is %q and my name is %q \n", age, nameOne)
	//q : if string then adds double quotes around else random garabage
	fmt.Printf("my age is %T and my name is %T \n", age, nameOne)
	//T : type of the variable
	fmt.Printf("my score is %f \n", scoretwo)
	fmt.Printf("my score is %0.2f \n", scoretwo)
	//f : float and decimal formating

	//Sprintf : it saves the string in var like input
	var strsave = fmt.Sprintf("my age is %v and my name is %v \n", age, nameOne)
	fmt.Println(strsave)

	// ARRAYS : ARRAY FIXED LENGTH only
	var ages [3]int = [3]int{10, 20, 30}
	names := [2]string{"hi", "he"}
	fmt.Println(ages, len(ages), names, len(names))

	//SLICES (uses arrays)
	var scores = []int{10, 40, 20}
	scores[2] = 25
	scores = append(scores, 85) //appends returns a new slice by def so u can override

	fmt.Println(scores, len(scores))

	//slice range
	rangestart := scores[1:3] // doesnt include 3 but inclusive 1
	rangefull := scores[1:]   // from 1 to end
	rangefull2 := scores[:3]  //till one before 3
	fmt.Println(rangestart, rangefull, rangefull2)

	//STANDARD LIBRARIES like fmt
	//GO TO MAIN2.GO

}

//to run : go run main.go
