package main

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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

var cookies = "session=53616c7465645f5f00da9b7b7e8a3118d029485d123f6a6bb8f0d3034220832735e03d0129a6e72acefc160c10d0805aa2ab817b8fb8507b436d17fd5fa8c5b5"

func main() {
	day := "11"

	file, err := os.Open("input" + day)

	content, err := io.ReadAll(file)
	file.Close()

	if err != nil {
		client := &http.Client{}

		req, err := http.NewRequest("GET", "https://adventofcode.com/2023/day/"+day+"/input", strings.NewReader(""))
		if err != nil {
			panic(err)
		}

		req.Header.Set("Cookie", cookies)
		resp, err := client.Do(req)

		content, err = io.ReadAll(resp.Body)

		os.WriteFile("input"+day, content, 0666)
	}

	out := mainDay11(string(content), 1000000)

	fmt.Println(Green + out + Reset)
}
