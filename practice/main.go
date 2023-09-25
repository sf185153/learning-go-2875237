package main

import "fmt"

func main() {
	ans := 42
	var result string

	if ans < 0 {
		result = "Less zero"
	} else if ans == 0 {
		result = "0"
	} else {
		result = "bigger than zero"
	}
	fmt.Println(result)

	if x := -42; x < 0 {
		result = "Less zero"
	} else if x == 0 {
		result = "0"
	} else {
		result = "bigger than zero"
	}
	fmt.Println(result)
}
