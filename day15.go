package main

import (
	"strconv"
	"strings"
)

func mainDay15(in string) string {
	sum := 0

	for _, s := range strings.Split(in, ",") {
		h := hash(s)
		sum += h
	}

	return strconv.Itoa(sum)
}

func hash(in string) int {
	current := 0

	for _, c := range in {
		// Determine the ASCII code for the current character of the string.
		// Increase the current value by the ASCII code you just determined.
		current += int(c)

		// Set the current value to itself multiplied by 17.
		current *= 17

		// Set the current value to the remainder of dividing itself by 256.
		current = current % 256
	}
	return current
}
