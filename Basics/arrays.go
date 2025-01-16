package main

import "fmt"

func getMessageWithRetries() [3]string {
	return [3] string {
		"Click here to sign up",
		"Pretty, please click here",
		"We beg you to sign up",
	}
}

func send(name string, doneAt int) {
	fmt.Printf("Sending to %v....", name)
	fmt.Println()

	messages := getMessageWithRetries()
	for i := 0; i < len(messages); i++ {
		msg:= messages[i]
		fmt.Printf(`Sending: "%v"`, msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("They Responded!!")
			break
		}
		if i == len(messages) - 1 {
			fmt.Println("They didn't respond")
		}
	}
}

func main() {
	send("John", 1)
	send("Jane", 2)
	send("Doe", 3)
	send("Smith", 4)
	send("Doe", 5)
}