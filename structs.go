package main

import "fmt"

func main() {
data := Dog{"German Sphered", 20, "barks"}
fmt.Println(data)
fmt.Printf("%+v\n", data)
fmt.Printf("Breed : %v , Weight : %v, Sound : %v\n", data.Breed, data.Weight, data.Sound)

fmt.Printf("\n")
data.DogBarks()

fmt.Printf("\n")
data.DogBarksThreeTimes()
}

type Dog struct {
	Breed string
	Weight int
	Sound string
}

func (d Dog) DogBarks() {
	fmt.Println("Dog Barks!!")
}

func (d Dog) DogBarksThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}