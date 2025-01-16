package main

import "fmt"

func main() {

	// Numbers are little different in Go
		// They fall under four categories
			// 1. Integers - Just whole numbers (-ve, +ve)
			// 2. Floating Point Numbers - Numbers with decimal points (fractional numbers)
			// 3. Complex Numbers - Numbers with real and imaginary parts
			// 4. Unsigned Integers - Just whole numbers (+ve)
	// Byte Type - It is an alias for uint8
	// Rune Type - It is an alias for int32

	// initializing variables
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}