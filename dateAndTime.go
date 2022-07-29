package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println("Today's date ->", n)
	t := time.Date(1998, time.February, 20, 20, 0, 0, 0, time.UTC)
	fmt.Println("UTC Date/Time -> ", t)
	fmt.Println("ANSIC Date/Time -> ", t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Fri Feb 20 20:00:00 1998")
	fmt.Printf("Type of ParsedTime -> %T\n", parsedTime)
}