package day1

import "fmt"

func q2(col1, col2 []int) {
	// Create frequency maps for each of the columns
	map1 := make(map[int]int)
	map2 := make(map[int]int)

	for i := 0; i < 1000; i++ {
		map1[col1[i]] = map1[col1[i]] + 1
		map2[col2[i]] = map2[col2[i]] + 1
	}

	similarityScore := 0
	for num, freq := range map1 {
		similarityScore += (num * map2[num] * freq)
	}

	fmt.Printf("Result for Day-1 Q2: %d\n", similarityScore)
}
