package main

import (
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	hand string
	bid  int
}

func mainDay7(in string) string {
	hands := Map(strings.Split(in, "\n"), func(line string) (h hand) {
		h.hand = strings.Split(line, " ")[0]
		bid, err := strconv.Atoi(strings.Split(line, " ")[1])
		if err != nil {
			panic("gjlrs")
		}
		h.bid = bid
		return
	})

	sort.Sort(hands)
}
