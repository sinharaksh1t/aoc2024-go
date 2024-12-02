package day1

import (
	"log"
	"strconv"
	"strings"

	"github.com/sinharaksh1t/aoc2024-go/helpers"
)

// Parse the input based on the requirement of the question
func parseInput() (col1, col2 []int) {
	body := helpers.FetchInput(FILE_NAME)

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
	return
}
