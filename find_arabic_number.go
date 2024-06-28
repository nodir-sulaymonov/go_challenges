package main

import (
	"fmt"
)

func main() {
    var str string
    fmt.Scan(&str)
	max := str[0]
    for i := range str {
		if max < str[i] {
			max = str[i]
		}
	}
	fmt.Print(string(max))

	//input 1112221112
}

