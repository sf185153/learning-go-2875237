package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("retype: %T", resp)

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	content := string(bytes)
	fmt.Println(content)
}
