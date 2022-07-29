package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func getInputValue() (float64){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter number")
	input1,_ := reader.ReadString('\n')
	n1,err1 := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err1 != nil {
		panic("Invalid value in input")
	}
	return n1
}

func main() {
    fmt.Println("#### Calculator App #####")
    reader := bufio.NewReader(os.Stdin)
	n1 := getInputValue()
	n2 := getInputValue()
	fmt.Println("Enter your choice")
	fmt.Println("Press 1 for Addition")
	fmt.Println("Press 2 for Substraction")
	fmt.Println("Press 3 for Multiplication")
	fmt.Println("Press 4 for Division")
	c,_ := reader.ReadString('\n')
	choice,err1 := strconv.ParseInt(strings.TrimSpace(c), 0, 64)
	if err1 != nil || choice != 1 && choice != 2 && choice != 3 && choice != 4  {
		panic("Invalid value in choice")
	} 

   var operation string
   var result float64 

	switch choice {
	case 1:
		operation = "Addition"
		result = n1+n2

	case 2:
		operation = "Substraction"
		result = n1-n2

	case 3:
		operation = "Multiplication"
		result = n1*n2

	case 4:
		operation = "Division"
		result = n1/n2

	default:
		fmt.Println("Invalid Choice")
	}
		
	fmt.Printf("By Performing %v operation on %v and %v result is %v", operation, n1, n2, result)
}
