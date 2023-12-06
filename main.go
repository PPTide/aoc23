package main

import (
	_ "embed"
	"fmt"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

//go:embed input6
var content string

func main() {
	out := mainDay6(content)

	fmt.Println(Green + out + Reset)
}
