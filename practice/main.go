package main

import (
	"fmt"
)

const aconst int = 64

func main() {

	var aString string = "Hello from Go!"
	fmt.Println(aString)
	fmt.Printf("the var is %T\n", aString)

	var inter int = 42
	fmt.Println(inter)
	fmt.Printf("the var is %T ", inter)

	var dfint int
	fmt.Println(dfint)

	var a2String = "String 2"
	fmt.Println(a2String)
	fmt.Printf("the var is %T\n", a2String)

	var int2 = 53
	fmt.Println(int2)
	fmt.Printf("the var is %T\n", int2)

	mystring := "Another string lmao"
	fmt.Println(mystring)
	fmt.Printf("the var is %T\n", mystring)
}
