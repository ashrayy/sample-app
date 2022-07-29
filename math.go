package main

import (
	"math"
	"fmt"
)

func main() {
	var floatNum = 3.1432
	var roundedNum = math.Round(floatNum)
	fmt.Printf("Orignal Number-> %v, Rounded-> %v\n", floatNum, roundedNum)

	n1, n2, n3 := 12.3, 3.456, 433.111
	sumFloat := n1+ n2 + n3
	fmt.Println("Sum without using math func->", sumFloat)

	sumFloat = math.Round(sumFloat*100)/100
	fmt.Println("Sum using math func->", sumFloat)

	circum := 34.5
	area := 2*(circum/2)*math.Pi
	fmt.Printf("Area -> %.2f\n", area)
}