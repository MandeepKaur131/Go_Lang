package main

import "fmt"

func bulkSend(numMsgs int) float64 {
	totalCost := 0.0
	for i := 0; i < numMsgs; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
	}
	return totalCost
}

func test(numMsgs int) {
	fmt.Printf("Sending %v messages\n", numMsgs)
	cost := bulkSend(numMsgs)
	fmt.Printf("Bulk sned complete! Cost = %.2f\n", cost)
}

func main() {
	test(3)
	test(10)
	test(7)
}