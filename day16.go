package main

import (
	"strconv"
	"strings"
)

var visited [][][]dir

func mainDay16(in string) string {
	lines := strings.Split(in, "\n")

	currentPos := pos{
		x: -1,
		y: 0,
	}
	currentDir := right

	visited = make([][][]dir, len(lines))

	for i := range visited {
		for j := 0; j < len(lines[0]); j++ {
			visited[i] = append(visited[i], make([]dir, 0))
		}
	}

	walkPath(currentPos, currentDir, lines)

	sum := 0

	for _, visits := range visited {
		for _, visit := range visits {
			if len(visit) > 0 {
				sum++
			}
		}
	}

	return strconv.Itoa(sum)
}

func walkPath(startPos pos, startDir dir, lines []string) {
	currentPos := startPos
	currentDir := startDir

	for {
		currentPos.x += dirToXYMap[currentDir].x
		currentPos.y += dirToXYMap[currentDir].y

		if currentPos.x < 0 || currentPos.x >= len(lines[0]) ||
			currentPos.y < 0 || currentPos.y >= len(lines) {
			break
		}

		if inSlice(currentDir, visited[currentPos.y][currentPos.x]) {
			break
		}

		visited[currentPos.y][currentPos.x] = append(visited[currentPos.y][currentPos.x], currentDir)

		switch lines[currentPos.y][currentPos.x] {
		case '.':
			continue
		case '/', '\\':
			if currentDir == right {
				currentDir = up
			} else if currentDir == up {
				currentDir = right
			} else if currentDir == left {
				currentDir = down
			} else if currentDir == down {
				currentDir = left
			}

			if lines[currentPos.y][currentPos.x] == '\\' {
				currentDir = opositeDirMap[currentDir]
			}
		case '|':
			if currentDir == left || currentDir == right {
				walkPath(currentPos, up, lines)
				walkPath(currentPos, down, lines)
				return
			}
		case '-':
			if currentDir == up || currentDir == down {
				walkPath(currentPos, left, lines)
				walkPath(currentPos, right, lines)
				return
			}
		}
	}
}
