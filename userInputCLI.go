package main

import (
	"fmt" // stands for format offers Printf and Println
	"bufio" //  package used for buffered IO
	"os" // provides a platform-independent interface to operating system functionality
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  // Note - \n at end makes user to input from next line
  fmt.Print("*** Enter Text *****\n")
  input,_ := reader.ReadString('\n')
  fmt.Println("You Entered -->", input)
}