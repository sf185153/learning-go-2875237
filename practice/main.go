package main

import (
	"fmt"
)

func main() {
	doSomething()
	sim := addVals(5, 8)
	fmt.Println(sim)
	fmt.Println(addAllVals(1, 2, 3, 4))
}

func doSomething() {
	fmt.Println("wowzers")
}

func addVals(val1 int, val2 int) int {
	return val1 + val2
}

func addAllVals(values ...int) (int, int) {
	total := 0
	for _, i := range values {
		total += i
	}
	return total, len(values)
}
