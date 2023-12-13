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

		lastRow := ""
		duplicateRowNums := make([]int, 0)
		for i, row := range rows {
			if row == lastRow {
				duplicateRowNums = append(duplicateRowNums, i-1)
			}
			lastRow = row
		}

		for _, num := range duplicateRowNums {
			if isMirrored(rows, num) {
				sumRow += num + 1
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

		lastCol := ""
		duplicateColNums := make([]int, 0)
		for i, col := range cols {
			if col == lastCol {
				duplicateColNums = append(duplicateColNums, i-1)
			}
			lastCol = col
		}

		for _, num := range duplicateColNums {
			if isMirrored(cols, num) {
				sumCol += num + 1
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

	for i := 0; i <= loopCount; i++ {
		if rows[upperIndex-i] != rows[upperIndex+1+i] {
			return false
		}
	}
	return true
}
