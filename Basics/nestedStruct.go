package main

import "fmt"

type messageToSend struct {
	message 		string
	sender 			user
	recipent 		user
}

type user struct {
	name 			string
	number 			string
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.recipent.name == "" {
		return false
	}
	if mToSend.recipent.number == "" {
		return false
	}
	if mToSend.sender.number == "" {
		return false
	}
	return true
}

func test(mToSend messageToSend) {
	if canSendMessage(mToSend) {
		fmt.Println("Message can be sent")
	} else {
		fmt.Println("Message cannot be sent")
	}
}

func main() {
	test(messageToSend{
		message: "Hello",
		sender: user{
			name: "John",
			number: "1234567890",
		},
		recipent: user{
			name: "Jane",
			number: "0987654321",
		},
	})
}