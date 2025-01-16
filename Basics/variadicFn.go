package main

import "fmt"

func main() {
	test(1.0, 2.0, 3.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0)
}

func sum(nums ...float64) float64 {
	sum := 0.0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum
}

func test(nums ...float64) {
	total := sum(nums...)
	fmt.Printf("Summing %v costs...\n", len(nums))
	fmt.Printf("Bill for the month: %.2f\n", total)
	fmt.Println("===== End Report =====")

}