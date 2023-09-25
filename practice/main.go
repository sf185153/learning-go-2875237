package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 15}
	fmt.Println(poodle.Breed)
}

// its a dog
type Dog struct {
	Breed  string
	Weight int
}
