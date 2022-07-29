package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello File Handling Works"
	file,err := os.Create("./file.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Println("Write a file with %v characters\n", length)
	defer file.Close()
	defer readFile("./file.txt")
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file is ->", string(data))
}

func checkError(err error){
	if err!= nil {
		panic(err)
	}
}