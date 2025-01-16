package main

import "fmt"

func main() {
	test(3, 3)
	test(4, 4)
	test(5, 5)
}

func createMatrix(rows, cols int) [][] int {
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		row := make([]int, 0)
		for j :=0; j < cols; j++ {
			row = append(row, i * j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func test(rows, cols int) {
	fmt.Printf("Creating a %d x %d matrix\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("==== End Report ====")
}