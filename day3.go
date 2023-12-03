package main

import (
	"strconv"
	"strings"
)

type number struct {
	num       int
	posxStart int
	posxEnd   int
	posy      int
}

func mainDay3(input string) string {
	lines := strings.Split(input, "\n")

	numbers := make([]number, 0)

	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			char := line[j]

			numString := ""

			for char >= '0' && char <= '9' {
				numString += string(char)
				j++
				if j >= len(line) {
					break
				}
				char = line[j]
			}

			if len(numString) > 0 {
				num, err := strconv.Atoi(numString)
				if err != nil {
					panic(err)
				}
				numbers = append(numbers, number{
					num:       num,
					posxStart: j - len(numString),
					posxEnd:   j,
					posy:      i,
				})
			}
		}
	}

	sum := 0

	for _, n := range numbers {
		if checkForSymbol(input, n) {
			sum += n.num
		}
	}

	return strconv.Itoa(sum)
}

func checkForSymbol(schematic string, n number) bool {
	lines := strings.Split(schematic, "\n")

	for x := n.posxStart - 1; x < n.posxEnd+1; x++ {
		if x < 1 || x >= len(lines[0]) { // is too far left or right?
			continue
		}

		y := n.posy - 1
		if y > 0 {
			char := lines[y][x]
			if (char < '0' || char > '9') && char != '.' {
				return true
			}
		}
		y = n.posy
		char := lines[y][x]
		if (char < '0' || char > '9') && char != '.' {
			return true
		}
		y = n.posy + 1
		if y <= len(lines)-2 {
			char := lines[y][x]
			if (char < '0' || char > '9') && char != '.' {
				return true
			}
		}
	}

	return false
}
