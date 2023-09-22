package main

import (
	"fmt"
)

func main() {
	anInt := 42
	var p = &anInt // p is a pointer, pointing at nil

	fmt.Println("P value, ", *p)
	*p = *p + 1
	fmt.Println("P value, ", *p)
	fmt.Println("Int is, ", anInt)
}
