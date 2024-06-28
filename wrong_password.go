package main

import (
	"fmt"
	"regexp"
)

func main() {
	var s1 string
	fmt.Scan(&s1)
	re := regexp.MustCompile( `^[a-zA-Z0-9]{5,}$`)
	if re.MatchString(s1){
		fmt.Print("Ok")
	} else {
		fmt.Print("Wrong password")
	}
	//input fdsghdfgjsdDD1
}