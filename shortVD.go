package main

import "fmt"

func main() {

	// A short way to declare and initialize variables
	// This is called short variable declaration
	// It is used when you want to declare and initialize a variable in one line
	// It is not necessary to specify the type of the variable
	// Go will automatically infer the type of the variable
	// It is a good practice to use short variable declaration when you know the
	// initial value of the variable
	// It is not a good practice to use short variable declaration when you don't
	// know the initial value of the variable
	// It is not a good practice to use short variable declaration when you want to
	// declare a variable and use it later in the program
	// It is not a good practice to use short variable declaration when you want to
	// declare a variable and assign a value to it later in the program
	// Example: var empty string is same as empty := ""
	// Example: var count int is same as count := 0
	// The short variable declartiion is what is used in Go most of the time

	congrats := "Congratulations! You have won a prize!"
	fmt.Println(congrats)
}