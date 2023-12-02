package main

import (
	"strconv"
	"strings"
)

var spelledOutDigits = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

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
		char := line[i]

		if char >= '0' && char <= '9' {
			num, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(err)
			}
			return num
		}
		for j, digit := range spelledOutDigits {
			if char == digit[len(digit)-1] {
				for k := len(digit) - 1; k >= 0; k-- {
					kCountingBack := len(digit) - k - 1
					if line[i-kCountingBack] != digit[k] {
						goto continueWithNextDigit
					}
				}
				return j
			}

		continueWithNextDigit:
			continue
		}
	}
	panic("No last digit found for: " + line)
}

func getFirstDigit(line string) int {
	for i := 0; i < len(line); i++ {
		char := line[i]

		if char >= '0' && char <= '9' {
			num, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(err)
			}
			return num
		}
		for j, digit := range spelledOutDigits {
			if char == digit[0] {
				for k, _ := range digit {
					if line[i+k] != digit[k] {
						goto continueWithNextDigit
					}
				}
				return j
			}

		continueWithNextDigit:
			continue
		}
	}
	panic("No first digit found for: " + line)
}
