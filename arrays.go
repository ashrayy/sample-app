package main

import (
	"fmt"
)

func main() {
	var colorsArr[3]string
	colorsArr[0] = "Red"
	colorsArr[1] = "Green"
	colorsArr[2] = "Blue"
	fmt.Println(colorsArr)
	fmt.Println(len(colorsArr))

	var numArr = [5]int{1,2,3,4,5}
	numArr[4] =6
	fmt.Println(numArr)
	fmt.Println(len(numArr))
}