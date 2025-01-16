package main

import "fmt"

type sender struct {
	rateLimit int
	user
}

type user struct {
	name 	string
	number 	int
}

func test(s sender) {
	fmt.Println("Sender name: ", s.name)
	fmt.Println("Sender number: ", s.number)
	fmt.Println("Sender rateLimit: ", s.rateLimit)
}

func main() {
	test(sender{
		rateLimit: 10,
		user: user{
			name: "John",
			number: 123,
		},
	})

	test(sender{
		rateLimit: 20,
		user: user{
			name: "Jane",
			number: 456,
		},
	})
}
