package day2

import "math"

func q2(report [][]int) int {
	safeReports := 0

	for _, row := range report {
		if inc, dec := checkSafe(row); inc || dec {
			safeReports++
		} else {
			for i, _ := range row {
				var newRow []int
				if i == 0 {
					newRow = row[1:]
				} else {
					newRow = append(newRow, row[:i]...)
					newRow = append(newRow, row[i+1:]...)
				}
				if inc, dec := checkSafe(newRow); inc || dec {
					safeReports++
					break
				}
			}
		}
	}

	return safeReports
}

func checkSafe(row []int) (bool, bool) {
	increasing := true
	decreasing := true
	for i := 0; i < len(row)-1; i++ {
		curr := row[i]
		next := row[i+1]
		if math.Abs(float64(curr-next)) > 3 {
			return false, false
		}
		if curr < next {
			decreasing = false
		} else if curr > next {
			increasing = false
		} else if curr == next {
			return false, false
		}

		if !increasing && !decreasing {
			return false, false
		}
	}
	return increasing, decreasing
}
