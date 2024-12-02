package day1

const (
	FILE_NAME = "day1/input.txt"
)

func Solve() {
	// Fetch input for this problem
	col1, col2 := parseInput()

	// Solve Problem 1
	q1(col1, col2)

	// Solve Problem 2
	q2(col1, col2)
}
