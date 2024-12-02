package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQ1(t *testing.T) {
	report := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
		{1, 1, 6, 7, 9},
		{1, 2, 9, 7, 9},
		{7, 6, 4, 2, 1, 0},
		{10, 7, 6, 4, 2, 1, 0},
		{11, 7, 6, 4, 2, 1, 0},
	}
	ansq1 := q1(report)

	assert.Equal(t, 4, ansq1)
}
