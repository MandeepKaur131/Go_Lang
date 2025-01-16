package main

import "fmt"

type messageToSend struct {
	phoneNumber string
	message     string
}

func test(m messageToSend) {
	fmt.Println("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("====================================")
}

func main() {
	test(messageToSend{
		phoneNumber: "1234567890",
		message:     "Thanks for signing up!",
	})

	test(messageToSend{
		phoneNumber: "0987654321",
		message:     "Thanks for signing up!",
	})

	test(messageToSend{
		phoneNumber: "1234567890",
		message:     "Thanks for signing up!",
	})

	test(messageToSend{
		phoneNumber: "0987654321",
		message:     "Thanks for signing up!",
	})
}
