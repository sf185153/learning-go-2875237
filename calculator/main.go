package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "hello"
	file, err := os.Create("./fromString.txt")
	checkErr(err)
	length, err := io.WriteString(file, content)
	checkErr(err)
	fmt.Printf("wrote a file with %v", length)
	defer file.Close()
	defer readFile("./fromString.txt")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkErr(err)
	fmt.Printf("text read %v, ", string(data))
}
