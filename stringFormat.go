package main

import "fmt"

func main() {
	const name = "John Doe"
	const openRate = 0.75
	msg := fmt.Sprintf("Hi %s, your open rate is %.1f", name, openRate)

	fmt.Println(msg)

}