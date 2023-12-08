package main

import (
	"fmt"
	"regexp"
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

	startNodes := make([]node, 0)

	var re = regexp.MustCompile(`(?m)(...) = \((...), (...)\)`)
	for _, line := range lines[2:] {
		matches := re.FindStringSubmatch(line)[1:]

		n := node{
			name:  matches[0],
			left:  matches[1],
			right: matches[2],
		}

		nodes[matches[0]] = n
		if n.name[2] == 'A' {
			startNodes = append(startNodes, n)
		}
	}

	lens := make([]int, 0)
	for _, n := range startNodes {
		i := 0
		for true {
			if instruction[i%len(instruction)] == 'L' {
				n = nodes[n.left]
			} else { // Fuck error handling xD
				n = nodes[n.right]
			}

			i++

			if n.name[2] == 'Z' {
				break
			}
		}
		lens = append(lens, i)
	}

	return "Please find the LCM: " + fmt.Sprint(lens)
}
