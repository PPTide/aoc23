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

	if len(dirs) != 2 {
		panic("Wrong start directions")
	}

	i := 0

	pos1 := start
	pos2 := start
	for {
		i++

		pos1.x += dirToXYMap[dirs[0]].x
		pos1.y += dirToXYMap[dirs[0]].y

		pos2.x += dirToXYMap[dirs[1]].x
		pos2.y += dirToXYMap[dirs[1]].y

		if pos1 == pos2 {
			break
		}

		dirs[0] = dirMap[lines[pos1.y][pos1.x]] ^ opositeDirMap[dirs[0]]
		dirs[1] = dirMap[lines[pos2.y][pos2.x]] ^ opositeDirMap[dirs[1]]
	}

	return strconv.Itoa(i)
}
