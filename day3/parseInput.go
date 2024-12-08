package day3

import (
	"strings"

	"github.com/sinharaksh1t/aoc2024-go/helpers"
)

// Parse the input based on the requirement of the question
func parseInput() []string {
	body := helpers.FetchInput(day, file_name)

	inputString := string(body)
	lines := strings.Split(inputString, "\n")

	return lines
}
