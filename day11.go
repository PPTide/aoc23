package main

import (
	"golang.org/x/exp/constraints"
	"strconv"
	"strings"
)

func mainDay11(in string) string {
	lines := strings.Split(in, "\n")

	newLines := make([]string, 0)

outer:
	for _, line := range lines {
		newLines = append(newLines, line)
		for _, c := range line {
			if c == '#' {
				continue outer
			}
		}
		newLines = append(newLines, line)
	}

	lines = newLines

outer2:
	for i := len(lines[0]) - 1; i >= 0; i-- {
		for _, line := range lines {
			if line[i] == '#' {
				continue outer2
			}
		}

		for j, line := range lines {
			newLine := make([]uint8, 0, len(lines[0]))

			newLine = append([]uint8(line[:i]), '.')
			newLine = append(newLine, []uint8(line[i:])...)

			lines[j] = string(newLine)
		}
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
			sum += Abs(position1.x - position2.x)
			sum += Abs(position1.y - position2.y)
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
