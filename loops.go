package main

import "fmt"

func main() {
	colorArr := []string{"Red","Green","Blue","Black","White"}

	for i := 0; i < len(colorArr); i++ {
		fmt.Println(colorArr[i])
	}
    
    fmt.Printf("\n")

	for i:= range colorArr {
		fmt.Println(colorArr[i])
	}

	fmt.Printf("\n")

	for i,color := range colorArr {
        fmt.Println(i,color)
	}

	fmt.Printf("\n")

	value:=0
	for value<10 {
		fmt.Println(value)
		value++
	}

	fmt.Printf("\n")

	sum:=0
	c:=0
	for c<20 {
		sum = sum+c
		if sum > 20 {
			goto theEnd
		}
		fmt.Println(sum)
		c++
	}

	theEnd: fmt.Println("Program End")
}