package day1

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func q1() {
	body, err := os.ReadFile(FILE_NAME)
	if err != nil {
		log.Fatalln("Error reading input file")
	}

	var col1, col2 []int
	inputString := string(body)
	lines := strings.Split(inputString, "\n")
	for _, line := range lines {
		input := strings.Fields(line)
		if len(input) == 0 {
			// Reached end of file
			break
		}
		if len(input) != 2 {
			log.Fatalln("Invalid input")
		}

		// Convert the values to int
		var1, err1 := strconv.Atoi(input[0])
		var2, err2 := strconv.Atoi(input[1])
		if err1 != nil || err2 != nil {
			log.Fatalln("Error parsing the value")
		}

		col1 = append(col1, var1)
		col2 = append(col2, var2)
	}

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
