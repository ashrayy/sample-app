package main

import "fmt"

func main() {
	var num = 121
	var ptrNum = &num

	fmt.Println("Number is ->", num)
	fmt.Println("Reference for Number is ->", ptrNum)
	fmt.Println("Number with ptr ref is ->", *ptrNum)

	*ptrNum = *ptrNum/12
    fmt.Println("Number is ->", num)

}