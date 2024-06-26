package main

import "fmt"

func main() {
	var n int
	cows := []string{"korov", "korova", "korovy", "korovy", "korovy","korov","korov","korov","korov","korov"}
	fmt.Scan(&n)

	if n > 0 && n < 100 {
		k := n
		if n > 10 && n < 20 {
			n = 0
		}
		n %= 10
		fmt.Printf("%d %s\n",k,cows[n])
	}
}