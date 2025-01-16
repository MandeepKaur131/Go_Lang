package main

import "fmt"

func main() {

	if length := 10; length < 5 {
		fmt.Println("Length is invalid.")
	} else {
		fmt.Println("Length is valid.")
	}
}