package main

import (
	"strconv"
	"strings"
)

func mainDay13(in string) string {
	parts := strings.Split(in, "\n\n")

	sumCol := 0
	sumRow := 0
outer:
	for _, part := range parts {
		rows := strings.Split(part, "\n")

		for i := 0; i < len(rows)-1; i++ {
			if isMirrored(rows, i) {
				sumRow += i + 1
				continue outer
			}
		}

		colsAsNum := make([][]int32, len(rows[0]))
		for _, row := range rows {
			for i, c := range row {
				colsAsNum[i] = append(colsAsNum[i], c)
			}
		}

		cols := Map(colsAsNum, func(t []int32) string {
			return string(t)
		})

		for i := 0; i < len(cols)-1; i++ {
			if isMirrored(cols, i) {
				sumCol += i + 1
				continue outer
			}
		}
	}

	return strconv.Itoa(sumCol + 100*sumRow)
}

func isMirrored(rows []string, upperIndex int) bool {
	loopCount := 0
	if upperIndex < len(rows)/2 {
		loopCount = upperIndex
	} else {
		loopCount = len(rows) - upperIndex - 2
	}

	smudges := 0

	for i := 0; i <= loopCount; i++ {
		if rows[upperIndex-i] != rows[upperIndex+1+i] {
			smudges += countDifferences(rows[upperIndex-i], rows[upperIndex+1+i])
			if smudges > 1 {
				return false
			}
		}
	}
	if smudges == 0 {
		return false
	}
	return true
}

func countDifferences(a, b string) (diffs int) {
	for i := range a {
		if a[i] != b[i] {
			diffs += 1
		}
	}
	return
}
