package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("Cannot divide by zero")
	}
	return x / y, nil
}

func test(x, y float64) {
	fmt.Printf("Dividing %.2f by %.2f\n", x, y)

	quotient, err := divide(x, y)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result: %.2f\n", quotient)
}

func main() {
	test(5.0, 0.0)
	test(5.0, 2.0)
}
