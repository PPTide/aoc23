package main

import (
	"strconv"
	"strings"
)

func mainDay1(input string) string {
	lines := strings.Split(input, "\n")
	nums := make([]int, 0, len(lines))

	for _, line := range lines {
		if len(line) < 1 {
			break
		}
		firstDigit := getFirstDigit(line)
		lastDigit := getLastDigit(line)

		num := firstDigit*10 + lastDigit
		nums = append(nums, num)
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	return strconv.Itoa(sum)
}

func getLastDigit(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= '0' && line[i] <= '9' {
			num, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(err)
			}
			return num
		}
	}
	panic("No last digit found for: " + line)
}

func getFirstDigit(line string) int {
	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			num, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(err)
			}
			return num
		}
	}
	panic("No first digit found for: " + line)
}
