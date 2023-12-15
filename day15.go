package main

import (
	"fmt"
	"strconv"
	"strings"
)

type lens struct {
	label    string
	focalLen int
}

func mainDay15(in string) string {
	boxes := make([][]lens, 256)

outer:
	for _, s := range strings.Split(in, ",") {
		//fmt.Println(boxes)
		if strings.ContainsRune(s, '=') {
			split := strings.Split(s, "=")
			label := split[0]
			fLen, _ := strconv.Atoi(split[1])
			l := lens{
				label:    split[0],
				focalLen: fLen,
			}
			for i, l2 := range boxes[hash(label)] {
				if l2.label == l.label {
					boxes[hash(label)][i] = l
					continue outer
				}
			}
			boxes[hash(label)] = append(boxes[hash(label)], l)
			continue outer
		}
		split := strings.Split(s, "-")
		label := split[0]
		for i, l := range boxes[hash(label)] {
			if l.label == label {
				//fmt.Println(label)
				//fmt.Println(boxes)
				boxes[hash(label)] = append(boxes[hash(label)][:i], boxes[hash(label)][i+1:]...)
			}
		}

	}

	fmt.Println(boxes)

	sum := 0

	for i, box := range boxes {
		for j, l := range box {
			sum += (i + 1) * (j + 1) * l.focalLen
		}
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
