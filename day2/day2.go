package day2

import "fmt"

const (
	FILE_NAME = "day2/input.txt"
)

func Solve() {
	// Fetch input for this problem
	report := parseInput()

	// Solve Problem 1
	ansq1 := q1(report)
	fmt.Printf("Result for Day-2 Q1: %d\n", ansq1)

	// Solve Problem 2
	ansq2 := q2(report)
	fmt.Printf("Result for Day-2 Q2: %d\n", ansq2)
}
