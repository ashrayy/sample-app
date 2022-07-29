package main

import (
	"fmt"
	"sort"
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

	var sliceableArr = []string{"Orange","Banana","Apple"}
	sliceableArr = append(sliceableArr,"Watermelon")
	fmt.Println(sliceableArr)
	fmt.Println(len(sliceableArr))

	sliceableArr = append(sliceableArr[1:len(sliceableArr)])
	fmt.Println(sliceableArr)
	fmt.Println(len(sliceableArr))

	sliceableArr = append(sliceableArr[:len(sliceableArr)-1])
	fmt.Println(sliceableArr)
	fmt.Println(len(sliceableArr))

	var numberArr = make([]int, 5)
	numberArr[0] = 12
	numberArr[1] = 243
    numberArr[2] = 886
	numberArr[3] = 7686
	numberArr[4] = 356
    fmt.Println(numberArr)

	sort.Ints(numberArr)
	fmt.Println(numberArr)
}