package main

import (
	"fmt"
)

func main() {
	var k, a, del int
	var arr []int
	fmt.Scan(&k)
	fmt.Scan(&del)
	for k > 0 {
		a = k % 10
		arr = append(arr, a)
		k /= 10
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != del {
			fmt.Print(arr[i])
		}
	}
}
