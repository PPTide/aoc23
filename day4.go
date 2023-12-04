package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func mainDay4(input string) string {
	lines := strings.Split(input, "\n")

	sum := 0
	for i, line := range lines {
		line = strings.Split(line, ":")[1]
		split := strings.Split(line, " |")

		var re = regexp.MustCompile(`(?m) (..)`)

		//FIXME: i am splitting empty thingies
		cardNums := re.FindAllString(split[0], -1)
		winningNums := re.FindAllString(split[1], -1)

		points := 0
		for _, num := range cardNums {
			if isInArray(winningNums, num) {
				if points > 0 {
					points *= 2
				} else {
					points = 1
				}
			}
		}

		fmt.Printf("Day %d: %d\n", i+1, points)

		sum += points
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
