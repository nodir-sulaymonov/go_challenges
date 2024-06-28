package main

import (
	"fmt"
)

func main() {
    var a string
    fmt.Scan(&a)
    for i, val := range a {
		if i == len(a) - 1 {
         fmt.Printf("%s", string(val))
			break
		}
		fmt.Printf("%s*", string(val))
	}
}