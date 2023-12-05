package main

import (
	"strconv"
	"strings"
)

func mainDay5(input string) string {
	categories := strings.Split(input, "\n\n")

	seedsString := categories[0]
	categories = categories[1:]

	seedsString = strings.Split(seedsString, ": ")[1]
	seeds := Map(strings.Split(seedsString, " "), func(t string) int {
		i, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		return i
	})

	for _, category := range categories {

	}
}

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}
