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

	timeStr := ""
	distStr := ""

	for _, time := range times {
		if len(time) > 0 {
			timeStr += time
		}
	}

	for _, dist := range dists {
		if len(dist) > 0 {
			distStr += dist
		}
	}

	time, err := strconv.Atoi(timeStr)
	if err != nil {
		panic(err)
	}
	dist, err := strconv.Atoi(distStr)
	if err != nil {
		panic(err)
	}

	posibilities := 0

	for j := 1; true; j++ {
		calcDist := calculateDistance(j, time)
		if calcDist <= dist && j < time {
			continue
		}
		if calcDist <= dist {
			break
		}
		//fmt.Printf("%d (dist: %d), ", j, dist)
		posibilities += 1
	}

	fmt.Printf(": %d\n", posibilities)

	return strconv.Itoa(posibilities)
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
