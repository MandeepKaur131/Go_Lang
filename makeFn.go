// Most of the times we don't bother ourselves abput underlying array of
// a slice. We can create a new slice using thr make function.
// Syntax: make([]T, len, cap)
// Here, T is the type of the slice to be created,
// len is the length of the slice
// and cap is the capacity of the slice.
// eg: s := make([]int, 5, 5)
// if we don't specify the capacity, it will be set to the
// length of the slice.
// eg: s := make([]int, 5)

package main

import "fmt"

func getMsgCosts(message [] string) [] float64 {

}

func test (messages [] string) {
	costs := getMsgCosts(messages)
	fmt.Println("Messages: ")
	for i := 0; i < len(messages); i++ {
		fmt.Printf("- %v\n", messages[i])
	}
	fmt.Println("Costs: ")
	for i := 0; i < len(costs); i++ {
		fmt.Printf("- %v\n", costs[i])
	}
}

func main() {
	test(getMsgCosts([]string{"hi", "hello", "bye"}))
	test(getMsgCosts([]string{"what's up", "see you"}))
}