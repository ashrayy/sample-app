package main

import (
	"fmt" 
	"bufio" 
	"os"
	"strconv"
	"strings"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("*** Enter A Number *****\n")
  input,_ := reader.ReadString('\n')
  fmt.Println("You Entered -->", input)
  numToFloat,err := strconv.ParseFloat(strings.TrimSpace(input), 64)
  if err!=nil {
	  fmt.Println(err)
  } else {
	  fmt.Println("Value of number", numToFloat)
	  fmt.Printf("Type is -> %T", numToFloat)
  }
}