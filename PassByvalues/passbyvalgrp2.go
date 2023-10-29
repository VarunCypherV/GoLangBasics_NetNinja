package main

import "fmt"

func updateMenu(x map[string]float64) {
	x["coffee"] = 2.99
}
func mainx() {
	menu := map[string]float64{
		"coffee":   3.011,
		"pie":      5.95,
		"icecream": 3.45,
	}
	updateMenu(menu)
	fmt.Println(menu)

}

//we get updated value and it actually alters the
//original value cuz unlike prev grp
// the data all is in a single block , pointer to data in other and more
// so bascially points and when u use the variable , go sees the pointer and
// the copy has a second pointer , a copy but essentially the same that also points to the same original block
//whereas in other grp copy are stored in seperate data blocks
//grp a non ptr values
//grp b pointer wrapper values
