package day1

import (
	"fmt"
	"math"
	"sort"
)

func q1() {
	col1, col2 := setupInput()

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

	fmt.Printf("Result for Day-1 Q1: %d\n", result)
}
