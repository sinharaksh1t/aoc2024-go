package day1

import "fmt"

const (
	FILE_NAME = "day1/input.txt"
)

func Solve() {
	// Fetch input for this problem
	col1, col2 := parseInput()

	// Solve Problem 1
	ansq1 := q1(col1, col2)
	fmt.Printf("Result for Day-1 Q1: %d\n", ansq1)

	// Solve Problem 2
	ansq2 := q2(col1, col2)
	fmt.Printf("Result for Day-1 Q2: %d\n", ansq2)
}
