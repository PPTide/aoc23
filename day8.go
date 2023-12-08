package main

import (
	"regexp"
	"strconv"
	"strings"
)

type node struct {
	name string

	left  string
	right string
}

func mainDay8(in string) string {
	lines := strings.Split(in, "\n")

	instruction := lines[0]
	nodes := make(map[string]node)

	var re = regexp.MustCompile(`(?m)(...) = \((...), (...)\)`)
	for _, line := range lines[2:] {
		matches := re.FindStringSubmatch(line)[1:]

		nodes[matches[0]] = node{
			name:  matches[0],
			left:  matches[1],
			right: matches[2],
		}
	}

	i := 0
	n := nodes["AAA"]
	for true {
		if instruction[i%len(instruction)] == 'L' {
			n = nodes[n.left]
		} else { // Fuck error handling xD
			n = nodes[n.right]
		}

		i++

		if n.name == "ZZZ" {
			break
		}
	}

	return strconv.Itoa(i)
}
