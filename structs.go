package main

import "fmt"

func main() {
data := Dog{"German Sphered", 20}
fmt.Println(data)
fmt.Printf("%+v\n", data)
fmt.Printf("Breed : %v , Weight : %v\n", data.Breed, data.Weight)
fmt.Println(data.Breed)
}

type Dog struct {
	Breed string
	Weight int
}