package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	fmt.Println(states)
	states["WA"] = "Washington"
	states["GA"] = "Georgia"
	states["CA"] = "California"
	fmt.Println(states)

	cali := states["CA"]
	fmt.Println(cali)

	delete(states, "WA")
	fmt.Println(states)

	states["NY"] = "New York"
	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
