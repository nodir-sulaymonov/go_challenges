package main

import (
	"fmt"
)
func main() {
	minimumFromFour()
}

func minimumFromFour() int {
	var n int
	var mint []int
	for i := 0; i < 4; i++ {
		fmt.Scan(&n)
		mint = append(mint, n)
	}
	var ml = mint[0]
	for i := 0; i < len(mint); i++ {
		if ml > mint[i] {
			ml = mint[i]
		}
	}
	return ml
    
}