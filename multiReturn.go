package main

import "fmt"

func main() {
	firstName, lastname := getFullName()
	fmt.Println("Welcome to Textio, ", firstName, lastname)

}

func getFullName() (string, string) {
	return "Sachin", "Tendulkar"
}