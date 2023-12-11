package main

import (
	"golang.org/x/exp/constraints"
	"strconv"
	"strings"
)

func mainDay11(in string, expansionVal int) string {
	lines := strings.Split(in, "\n")

	lineHeights := make([]int, 0, len(lines))

outer:
	for _, line := range lines {
		for _, c := range line {
			if c == '#' {
				lineHeights = append(lineHeights, 1)
				continue outer
			}
		}
		lineHeights = append(lineHeights, expansionVal)
	}

	colWidths := make([]int, 0, len(lines[0]))

outer2:
	for i := range lines[0] {
		for _, line := range lines {
			if line[i] == '#' {
				colWidths = append(colWidths, 1)
				continue outer2
			}
		}

		colWidths = append(colWidths, expansionVal)
	}

	// --------------- Calculate all distances -------------------
	positions := make([]pos, 0)

	for i, line := range lines {
		for j, c := range line {
			if c == '#' {
				positions = append(positions, pos{
					x: j,
					y: i,
				})
			}
		}
	}

	sum := 0

	for i, position1 := range positions {
		for _, position2 := range positions[i:] {
			for _, i2 := range sliceSlice(lineHeights, position1.y, position2.y) {
				sum += i2
			}
			for _, i2 := range sliceSlice(colWidths, position1.x, position2.x) {
				sum += i2
			}
		}
	}

	return strconv.Itoa(sum)
}

func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func sliceSlice[T any](s []T, a, b int) []T {
	if a < b {
		return s[a:b]
	}
	return s[b:a]
}
