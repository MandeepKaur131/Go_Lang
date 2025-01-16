package main

import "fmt"

// func concat(a string, b string) string {
// 	return a + b
// }

func concat(a, b string) string {
	return a + b
}

func main() {
	fmt.Println(concat("Hello, ", "World!"))
	fmt.Println(concat("Go, ", "Went, Gone!"))
}
