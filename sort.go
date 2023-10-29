package main

import (
	"fmt"
	"sort"
)

func main3() {
	ages := []int{1, 6, 3, 4, 5}
	sort.Ints(ages) //for sorting based on int
	//This Method alter original slice
	index := sort.SearchInts(ages, 5)
	// If element not found it returns index of last element + 1
	fmt.Println(ages, index)
	names := []string{"raj", "ram", "rakesh"}
	sort.Strings(names)
	indexs := sort.SearchStrings(names, "raj")
	fmt.Println(indexs)
	//same for float
}
