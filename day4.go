package main

import (
	"regexp"
	"strconv"
	"strings"
)

func mainDay4(input string) string {
	lines := strings.Split(input, "\n")

	numOfCards := make([]int, 0, len(lines))
	for i := 0; i < len(lines); i++ {
		numOfCards = append(numOfCards, 1)
	}

	for i, line := range lines {
		line = strings.Split(line, ":")[1]
		split := strings.Split(line, " |")

		var re = regexp.MustCompile(`(?m) (..)`)

		cardNums := re.FindAllString(split[0], -1)
		winningNums := re.FindAllString(split[1], -1)

		matches := 0
		for _, num := range cardNums {
			if isInArray(winningNums, num) {
				matches += 1
			}
		}

		for j := i; j < i+matches; j++ {
			numOfCards[j+1] += numOfCards[i]
		}
	}

	sum := 0
	for _, n := range numOfCards {
		sum += n
	}

	return strconv.Itoa(sum)
}

func isInArray[T comparable](array []T, element T) bool {
	for _, t := range array {
		if element == t {
			return true
		}
	}
	return false
}
