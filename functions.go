package main

import (
	"fmt"
)

func main() {
  doSomething()
  fmt.Println("Sum is :-", addValues(2,3))
  sum, length := addAllValues(2,3,4,5,6,7,8,9)
  fmt.Println("Sum is :-", sum)
  fmt.Println("Number of values are :-", length)
}

func doSomething() {
	fmt.Println("Doing Something")
}

func addValues(value1 int, value2 int) int{
	return value1 + value2
}

func addAllValues(values ...int) (int,int) {
	total := 0
	for _,v := range values {
         total+=v
	}
	return total, len(values)
}