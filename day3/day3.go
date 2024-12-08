package day3

import "fmt"

const (
	day       = 3
	file_name = "day3/input.txt"
)

func Solve() {
	// Fetch input for this problem
	input := parseInput()

	// Solve Problem 1
	ansq1 := q1(input)
	fmt.Printf("Result for Day-3 Q1: %d\n", ansq1)

	// Solve Problem 2
	// ansq2 := q2(report)
	// fmt.Printf("Result for Day-2 Q2: %d\n", ansq2)
}
