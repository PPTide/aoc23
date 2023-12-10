package main

import (
	"strconv"
	"strings"
)

type dir int8

const (
	up dir = 1 << iota
	right
	down
	left
)

var dirMap = map[uint8]dir{
	'F': right | down,
	'J': up | left,
	'L': up | right,
	'7': left | down,
	'|': up | down,
	'-': left | right,
	'.': 0,
}

var reverseDirMap = reverseMap(dirMap)

var dirToXYMap = map[dir]pos{
	up:    {y: -1},
	down:  {y: 1},
	right: {x: 1},
	left:  {x: -1},
}

var opositeDirMap = map[dir]dir{
	up:    down,
	down:  up,
	right: left,
	left:  right,
}

type pos struct {
	x, y int
}

func mainDay10(in string) string {
	lines := strings.Split(in, "\n")

	var start pos

outer:
	for y, line := range lines {
		for x, c := range line {
			if c == 'S' {
				start.y = y
				start.x = x
				break outer
			}
		}
	}

	dirs := make([]dir, 0, 2)

	for d, relativePos := range dirToXYMap {
		if start.x+relativePos.x >= len(lines[0]) || start.x+relativePos.x < 0 ||
			start.y+relativePos.y >= len(lines) || start.y+relativePos.y < 0 {
			continue
		}

		if dirMap[lines[start.y+relativePos.y][start.x+relativePos.x]]&opositeDirMap[d] != 0 {
			dirs = append(dirs, d)
		}
	}

	line := assignToSlice([]uint8(lines[start.y]), reverseDirMap[dirs[0]|dirs[1]], start.x)

	lines = assignToSlice(lines, string(line), start.y)

	if len(dirs) != 2 {
		panic("Wrong start directions")
	}

	i := 0

	pos1 := start
	pos2 := start

	loop := make([]pos, 0, 6956*2) // I know the loop size from part1
	for {
		i++

		loop = append(loop, pos1)
		loop = append(loop, pos2)

		pos1.x += dirToXYMap[dirs[0]].x
		pos1.y += dirToXYMap[dirs[0]].y

		pos2.x += dirToXYMap[dirs[1]].x
		pos2.y += dirToXYMap[dirs[1]].y

		if pos1 == pos2 {
			loop = append(loop, pos1)
			break
		}

		dirs[0] = dirMap[lines[pos1.y][pos1.x]] ^ opositeDirMap[dirs[0]]
		dirs[1] = dirMap[lines[pos2.y][pos2.x]] ^ opositeDirMap[dirs[1]]
	}

	i = 0
	for y, line := range lines {
		for x, c := range line {
			p := pos{x: x, y: y}

			if !inSlice(p, loop) {
				if inLoop(p, loop, lines) {
					i++
					if !inLoopFromRight(p, loop, lines) {
						print(Red)
					}
					print("I", Reset)
				} else {
					if inLoopFromRight(p, loop, lines) {
						print(Red)
					}
					print("O", Reset)
				}
			} else {
				if inLoop(p, loop, lines) {
					//print(Blue)
				}
				print(string(niceChar(c)), Reset)
			}
		}
		println()
	}

	return strconv.Itoa(i)
}

func niceChar(c rune) rune {
	switch c {
	case 'F':
		return '┏'
	case 'L':
		return '┗'
	case 'J':
		return '┛'
	case '7':
		return '┓'
	case '-':
		return '━'
	case '|':
		return '┃'
	default:
		return c
	}
}

func inSlice[T comparable](a T, list []T) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func inLoop(p pos, loop []pos, pipeMap []string) bool {
	// Look from the bottom if there is an even or odd number of pipes in front of p
	look := p
	look.y = len(pipeMap) - 1

	cutOvers := 0
	for ; look.y > p.y; look.y-- {
		if isInArray(loop, look) {
			c := pipeMap[look.y][look.x]

			switch c {
			case '-':
				cutOvers += 2 // plus one
			case 'F', 'J':
				cutOvers += 1 // plus one half
			case '7', 'L':
				cutOvers -= 1 // minus one half
			}
		}
	}

	if cutOvers%2 == 1 {
		//panic("WTF")
	}

	return (cutOvers)%4 == 2
}

func inLoopFromRight(p pos, loop []pos, pipeMap []string) bool {
	// Look from the bottom if there is an even or odd number of pipes in front of p
	look := p
	look.x = len(pipeMap[0]) - 1

	cutOvers := 0
	for ; look.x > p.x; look.x-- {
		if isInArray(loop, look) {
			c := pipeMap[look.y][look.x]

			switch c {
			case '|':
				cutOvers += 2 // plus one
			case 'F', 'J':
				cutOvers += 1 // plus one half
			case '7', 'L':
				cutOvers -= 1 // minus one half
			}
		}
	}

	if cutOvers%2 == 1 {
		panic("WTF")
	}

	return (cutOvers)%4 == 2
}

func reverseMap[T, U comparable](m map[T]U) map[U]T {
	n := make(map[U]T, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

func assignToSlice[T any](s []T, e T, pos int) []T {
	n := make([]T, 0, len(s))
	for i, t := range s {
		if i == pos {
			n = append(n, e)
			continue
		}
		n = append(n, t)
	}

	return n
}
