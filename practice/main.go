package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)
	colors = append(colors[1 : len(colors)-1])
	fmt.Println(colors)

	sort.Strings(colors)
	fmt.Println(colors)
}
