package main

import (
	"fmt"
	"strconv"
	"strings"
)

func mainDay6(in string) string {
	lines := strings.Split(in, "\n")
	times := strings.Split(lines[0], " ")[1:]
	dists := strings.Split(lines[1], " ")[1:]

	ts := Map(removeEmptyStringFromSlice(times), func(t string) int {
		i, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		return i
	})
	ds := Map(removeEmptyStringFromSlice(dists), func(t string) int {
		i, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		return i
	})

	if len(ts) != len(ds) {
		panic("not the same amount of times and distances")
	}

	prod := 1

	for i, time := range ts {
		posibilities := 0

		for j := 1; true; j++ {
			dist := calculateDistance(j, time)
			if dist <= ds[i] && j < time {
				continue
			}
			if dist <= ds[i] {
				break
			}
			fmt.Printf("%d (dist: %d), ", j, dist)
			posibilities += 1
		}

		fmt.Printf(": %d\n", posibilities)
		prod *= posibilities
	}

	return strconv.Itoa(prod)
}

func calculateDistance(timeHeld int, maxTime int) int {
	restTime := maxTime - timeHeld

	return restTime * timeHeld // restTime * speed
}

func removeEmptyStringFromSlice(in []string) (out []string) {
	for _, s := range in {
		if len(s) > 0 {
			out = append(out, s)
		}
	}
	return
}
