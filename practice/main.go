package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("int sum is: ", intSum)

	f1, f2, f3 := 23.5, 62.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("float sum is: ", floatSum)

	floatSum = math.Round(floatSum)
}
