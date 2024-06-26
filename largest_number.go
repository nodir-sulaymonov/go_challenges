package main

import (
	"fmt"
)

func main() {
	var a, b, maxM, count int
	fmt.Scan(&a)
	fmt.Scan(&b)
	maxM = a
	for i := a; i <= b; i++ {
		if maxM < b && i%7 == 0 {
			maxM = i
			count++
		}
	}

	if count > 0 {
		fmt.Print(maxM)
	} else {
		fmt.Print("NO")
	}

}
