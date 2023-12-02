package main

import (
	"regexp"
	"strconv"
	"strings"
)

type numCubes struct {
	num   int
	color string
}

type set []numCubes

type game struct {
	sets  []set
	index int
}

func mainDay2(input string) string {
	var generalStructure = regexp.MustCompile(`(?m)Game (\d+): (.+)`)
	var splitColorNum = regexp.MustCompile(`(?m)(\d+) (.+)`)
	games := make([]game, 0)

	for _, line := range strings.Split(input, "\n") {
		if len(line) < 1 {
			break
		}
		match := generalStructure.FindStringSubmatch(line)

		thisGame := game{}

		index, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}

		thisGame.index = index

		record := match[2]
		sets := make([]set, 0)
		for _, setString := range strings.Split(record, "; ") {
			thisSet := make(set, 0)
			for _, color := range strings.Split(setString, ", ") {
				match := splitColorNum.FindStringSubmatch(color)
				num, err := strconv.Atoi(match[1])
				if err != nil {
					panic(err)
				}
				thisSet = append(thisSet, numCubes{
					num:   num,
					color: match[2],
				})
			}
			sets = append(sets, thisSet)
		}

		thisGame.sets = sets

		games = append(games, thisGame)
	}

	res := 0

	for _, game := range games {
		maxRed := 1
		maxGreen := 1
		maxBlue := 1
		for _, set := range game.sets {
			for _, cubes := range set {
				switch cubes.color {
				case "red":
					if cubes.num > maxRed {
						maxRed = cubes.num
					}
				case "green":
					if cubes.num > maxGreen {
						maxGreen = cubes.num
					}
				case "blue":
					if cubes.num > maxBlue {
						maxBlue = cubes.num
					}
				}
			}
		}

		res += maxRed * maxGreen * maxBlue
	}

	return strconv.Itoa(res)
}
