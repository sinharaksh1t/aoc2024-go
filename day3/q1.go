package day3

import (
	"regexp"
	"strconv"
	"strings"
)

func q1(input []string) int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	result := 0
	for _, line := range input {
		matches := re.FindAll([]byte(line), -1)
		for _, match := range matches {
			str := string(match)
			numSplit := strings.Split(str[4:len(str)-1], ",")
			num1, _ := strconv.Atoi(numSplit[0])
			num2, _ := strconv.Atoi(numSplit[1])
			result += (num1 * num2)
		}
	}

	return result
}
