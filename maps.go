package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	fmt.Println(states)
	states["UP"] = "Uttar Pradesh"
	states["UK"] = "Uttrakhand"
	states["DEL"] = "Delhi"
	states["RJ"] = "Rajasthan"
	fmt.Println(states)

	stateName := states["DEL"]
    fmt.Println(stateName)

	delete(states,"RJ")
	fmt.Println(states)

	for k,v := range states {
		fmt.Printf("%v : %v\n", k, v)
	}

	statesKeyArr := make([]string, len(states))
	i := 0
	for k := range states {
		statesKeyArr[i] = k
		i++
	}

	fmt.Println(statesKeyArr)
	sort.Strings(statesKeyArr)
	fmt.Println(statesKeyArr)

	for i := range statesKeyArr {
		fmt.Println(states[statesKeyArr[i]])
	} 
	
}