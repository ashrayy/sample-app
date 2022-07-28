package main

import "fmt" // stands for format offers Printf and Println

const tempConst = 100

func main() {
	var tempString string = "This is String"
    fmt.Println(tempString)
	fmt.Printf("Variable Type -> %T\n", tempString)

	var tempInteger int = 11111
	fmt.Println(tempInteger)
	fmt.Printf("Variable Type -> %T\n", tempInteger)

	var stringWithoutType = "This is String 2"
	fmt.Println(stringWithoutType)
	fmt.Printf("Variable Type for string Without Type -> %T\n", stringWithoutType)
    
	// NOTE - works only inside function
	newString:= "This is String 3"
	fmt.Println(newString)
	fmt.Printf("Variable Type for string With Colon -> %T\n", newString)

	// NOTE - accessing global variables
	fmt.Println(tempConst)
	fmt.Printf("Variable Type for declared constant -> %T\n", tempConst)


}