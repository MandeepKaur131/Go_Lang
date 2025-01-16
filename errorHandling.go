package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	costToCustomer, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	costToSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0.0, err
	}
	return costToCustomer + costToSpouse, nil
}

func sendSMS(msg string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = .0002
	if len(msg) > maxTextLen {
		return 0, fmt.Errorf("Message is too long")
	}
	return float64(len(msg)) * costPerChar, nil
}

func test(msgToCustomer, msgToSpouse string) {
	defer fmt.Println("====================")
	fmt.Println("Message for customer: ", msgToCustomer)
	fmt.Println("Message for spouse: ", msgToSpouse)
	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("Total cost: $%.4f\n", totalCost)
}

func main() {
	message := "Hello"
	test(message, message)
	m1 := "Hello, I am going to the store to buy some groceries."
	m2 := "let's meet at the park"
	test(m1, m2)
}