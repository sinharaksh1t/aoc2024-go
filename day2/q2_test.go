package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQ2Safe(t *testing.T) {
	report := [][]int{
		// Safe
		{7, 6, 4, 2, 1},
		{7, 6, 4, 2, 2},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{8, 6, 4, 5, 1},
		{1, 3, 6, 7, 9},
		{1, 5, 6, 7, 9},
		{1, 2, 9, 3, 4},
		{7, 6, 4, 2, 1, 0},
		{10, 7, 6, 4, 2, 1, 0},
		{11, 7, 6, 4, 2, 1, 0},
	}
	ansq2 := q2(report)

	assert.Equal(t, len(report), ansq2)
}

func TestQ2Unsafe(t *testing.T) {
	report := [][]int{
		// Unsafe
		{9, 7, 6, 8, 2, 1, 0},
		{1, 2, 7, 8, 9},
		{1, 2, 9, 7, 8},
		{9, 7, 6, 2, 1},
		{1, 1, 6, 7, 9},
		{11, 7, 6, 4, 2, 1, 1},
		{1, 1, 1, 1},
		{1, 2, 3, 3, 2, 1},
		{8, 6, 4, 5, 1, 0, 1, 0},
	}
	ansq2 := q2(report)

	assert.Equal(t, 0, ansq2)
}
