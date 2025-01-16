package main

import "fmt"

func main() {
	accountAge := 2.6

	// Create a new "accountAgeInt" here
	// It should be the result of converting/casting "accountAge" to an integer
	accountAgeInt := int(accountAge)

	fmt.Println("Your account has existed for", accountAgeInt, "years")
}