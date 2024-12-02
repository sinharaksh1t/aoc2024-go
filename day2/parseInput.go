package day2

import (
	"log"
	"strconv"
	"strings"

	"github.com/sinharaksh1t/aoc2024-go/helpers"
)

// Parse the input based on the requirement of the question
func parseInput() [][]int {
	body := helpers.FetchInput(2, FILE_NAME)

	var report [][]int
	inputString := string(body)
	lines := strings.Split(inputString, "\n")
	for _, line := range lines {
		input := strings.Fields(line)
		if len(input) == 0 {
			// Reached end of file
			break
		}

		var row []int
		for _, val := range input {
			num, err := strconv.Atoi(val)
			if err != nil {
				log.Fatalln("Error parsing the value")
			}
			row = append(row, num)
		}
		report = append(report, row)
	}

	return report
}
