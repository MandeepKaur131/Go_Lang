package main

import "fmt"

func getFormattedMessages(messages [] string, formatter func (string) string) [] string {
	formattedMessages := [] string {}
	for _, message := range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}
	return formattedMessages
}

func addSignature(message string) string {
	return message + " Kind regards."
}

func addGreeting(message string) string {
	return "Hello! " + message
}

func test(message [] string, formatter func(string) string) {
	
}

func main() {}