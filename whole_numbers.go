package main

import (
	"fmt"
	"math"
)
func main() {
	var a, degree float64
	var slice []int
	fmt.Scan(&a)
	degree = 0
	for {
		square := math.Pow(2, degree)
		if square <= a {
			squarex := int(square)
			slice = append(slice, squarex)
			degree++
		} else if square > a {
			for i := 0; i < len(slice); i++ {
				fmt.Printf("%d ", slice[i])
			}
			break
		}
	}
}
