package main

import (
	"fmt"
	"time"
)

func main() {

	n := time.Now()
	fmt.Println("Dates and times, ", n)

	t := time.Date(2009, time.April, 23, 4, 1, 50, 0, time.UTC)
	fmt.Println("made date: ", t)

	parseTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")

	fmt.Printf(parseTime.String())
}
