package day1

import (
	"math"
	"sort"
)

func q1(col1, col2 []int) int {
	// Sort the two columns
	// I know O(n log(n)), eww
	sort.Ints(col1)
	sort.Ints(col2)

	var result int = 0
	for i := 0; i < 1000; i++ {
		val1 := col1[i]
		val2 := col2[i]

		result += int(math.Abs(float64(val1 - val2)))
	}

	return result
}
