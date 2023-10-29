package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bill struct { //no need commas
	name  string
	items map[string]float64
	tip   float64
}

// recevier func : associating this func with bill struct
// we r limiting that this func only cud be called by a bill object
func (b bill) format() string {
	//now say we pass my bill , this b here is mybill
	fs := "Bill BreakDown : \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ....$%v \n", k+":", v)
		//-25 : it takes 25 characters  if not there then space : for allignment
		//if +25 then it takes space towards left , else right

		total += v
	}
	fs += fmt.Sprintf("%-25v ....$%0.2f \n", "tip:", b.tip)
	fs += fmt.Sprintf("%-25v ....$%0.2f \n", "total:", total+b.tip)
	return fs
}

// we pass ptr since func itself is grp 1 and makes copy and also
// we reduce the number of copies of bill
func (b *bill) updateTip(tip float64) {
	// (*b).tip = tip  need to do
	//but when struct , go auto derefs when struct
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save() {
	//to save in file we need to save in slice in byte
	//WRITE FILE
	data := []byte(b.format())

	//0644 : permission
	err := os.WriteFile("billForStruct1/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err) //stops flow of prog and terminates : os
	}
	fmt.Println("Saved to folder")
}

// returning bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}
	return b
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	// //reader reads from standard input , name stores everything until newline
	name, _ := getInput("Create a new Bill Name: ", reader)
	b := newBill(name)
	fmt.Println("Created New Bill - ", b.name)
	return b
}

// accepting the reader for taking input instead of declare everywhere
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	option, _ := getInput("Choose option a-add item \ns- save bill\nt- add tip\n choice :  ", reader)

	// SWITCH CASE IN GO LAB
	switch option {
	case "a":
		name, _ := getInput("Item name : ", reader)
		price, _ := getInput("Item Price : ", reader)
		//PARSING IN GO LANG
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Price Must be a Number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Added!")
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("Saved as ", b.name)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Tip : ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Price Must be a Number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Added!")
		promptOptions(b)
		fmt.Println(tip)
	default:
		fmt.Println("unknown")
		promptOptions(b)
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
	mybill.addItem("onion soup", 4.50)
	mybill.updateTip(10)

}

//NOTE FOR TUTORIAL PURPOSE I DID POINTER AND NON POINTER BOTH
//BEST PRACTICE IS TO USE ONE EVERYWHERE
