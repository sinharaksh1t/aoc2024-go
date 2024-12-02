package day2

import (
	"math"
)

func q1(report [][]int) int {
	safeReports := 0

	for _, row := range report {
		val1 := row[0]
		val2 := row[1]
		diff := math.Abs(float64(val1 - val2))

		if val1 < val2 && diff <= 3 {
			safeReports += checkIncreasing(row)
		} else if val1 > val2 && diff <= 3 {
			safeReports += checkDecreasing(row)
		}
	}

	return safeReports
}

func checkIncreasing(row []int) int {
	for i := 1; i < len(row)-1; i++ {
		curr := row[i]
		next := row[i+1]

		if !(curr < next && math.Abs(float64(curr-next)) <= 3) {
			return 0
		}
	}
	return 1
}

func checkDecreasing(row []int) int {
	for i := 1; i < len(row)-1; i++ {
		curr := row[i]
		next := row[i+1]

		if !(curr > next && math.Abs(float64(curr-next)) <= 3) {
			return 0
		}
	}
	return 1
}
