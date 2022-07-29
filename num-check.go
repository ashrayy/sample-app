package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main() {
    fmt.Println("#### Value Checker App #####")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter value")
	input1,_ := reader.ReadString('\n')
	n1,err1 := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err1 != nil {
		panic("Invalid value . Please input numeric value")
	} 

    if n1 > 0 {
		fmt.Println("Positive value")
	} else if n1 < 0 {
		fmt.Println("Negative value")
	} else {
		fmt.Println("Zero")
	}

	// assignment with if statement
	if x := -42; x < 0 {
		fmt.Println("Less")
	} else if x > 0 {
		fmt.Println("More")
	} else {
		fmt.Println("Zero")
	}
}