package main

import "fmt"

func main() {
	msgLen := 10
	maxMsgLen := 20
	fmt.Println("Trying to send a message of length:", msgLen)

	if msgLen > maxMsgLen {
		fmt.Println("Message too long!")
	} else {
		fmt.Println("Message sent successfully!")
	}
}
