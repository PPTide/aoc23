package main

import (
	"math"
	"strconv"
	"strings"
)

type intRange struct {
	sourceStart int64
	destStart   int64
	length      int64
}

func mainDay5(input string) string {
	categories := strings.Split(input, "\n\n")

	seedsString := categories[0]
	categories = categories[1:]

	seedsString = strings.Split(seedsString, ": ")[1]
	seeds := Map(strings.Split(seedsString, " "), func(t string) int64 {
		i, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			panic(err)
		}
		return i
	})

	newSeeds := make([]int64, 0, len(seeds)/2)
	for i := 0; i < len(seeds); i += 2 {
		for j := int64(0); j < seeds[i+1]; j++ {
			newSeeds = append(newSeeds, seeds[i]+j)
		}
	}

	seeds = newSeeds

	for _, category := range categories {
		ranges := Map(strings.Split(category, "\n")[1:], func(r string) intRange {
			nums := Map(strings.Split(r, " "), func(t string) int64 {
				i, err := strconv.ParseInt(t, 10, 64)
				if err != nil {
					panic(err)
				}
				return i
			})
			return intRange{
				sourceStart: nums[1],
				destStart:   nums[0],
				length:      nums[2],
			}
		})

		for i, seed := range seeds {
			for _, r := range ranges {
				if seed >= r.sourceStart && seed <= r.sourceStart+r.length {
					afterStart := seed - r.sourceStart

					seeds[i] = r.destStart + afterStart

					goto continueWithNextSeed
				}
			}
			// There is no match so the number stays the same

		continueWithNextSeed:
			continue
		}
	}

	var minLoc int64 = math.MaxInt64
	for _, seed := range seeds {
		if seed < minLoc {
			minLoc = seed
		}
	}

	return strconv.FormatInt(minLoc, 10)
}

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}
