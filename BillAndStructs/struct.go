package main

import "fmt"

type bill struct { //no need commas
	name  string
	items map[string]float64
	tip   float64
}

//recevier func : associating this func with bill struct
//we r limiting that this func only cud be called by a bill object
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

//we pass ptr since func itself is grp 1 and makes copy and also
//we reduce the number of copies of bill
func (b *bill) updateTip(tip float64) {
	// (*b).tip = tip  need to do
	//but when struct , go auto derefs when struct
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
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

func mainx() {
	mybill := newBill("Mario Bill")
	mybill.addItem("onion soup", 4.50)
	mybill.updateTip(10)
	fmt.Println(mybill.format())
}

//NOTE FOR TUTORIAL PURPOSE I DID POINTER AND NON POINTER BOTH
//BEST PRACTICE IS TO USE ONE EVERYWHERE
