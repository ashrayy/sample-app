package main

import (
	"fmt"
	"os"
	"math"
	"bufio"
	"strconv"
	"strings"
)

func main() {
    fmt.Println("#### Calculator App #####")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter number 1")
	input1,_ := reader.ReadString('\n')
	n1,err1 := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err1 != nil {
		panic("Invalid value in input1")
	} 
	fmt.Println("Enter number 2")
	input2,_ := reader.ReadString('\n')
	n2,err2 := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err2 != nil {
		panic("Invalid value in input2")
	} 
	sum := n1+n2
	fmt.Printf("Sum of number1 is %v , number2 is %v and sum-> %v\n", n1, n2, math.Round(sum))
}
