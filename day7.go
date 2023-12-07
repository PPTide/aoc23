package main

import (
	"fmt"
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

	// sort by custom func
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].value() == hands[j].value() {
			for k := 0; k < len(hands[i].hand); k++ {
				if cardToValue(rune(hands[i].hand[k])) == cardToValue(rune(hands[j].hand[k])) {
					continue
				}
				return cardToValue(rune(hands[i].hand[k])) < cardToValue(rune(hands[j].hand[k]))
			}
		}
		return hands[i].value() < hands[j].value()
	})

	sum := 0

	for i, h := range hands {
		sum += (i + 1) * h.bid
	}

	return strconv.Itoa(sum)
}

func (h hand) value() int {
	cardAmounts := make([]int, 15)
	jokers := 0

	for _, c := range h.hand {
		if c == 'J' {
			jokers++
			continue
		}
		cardAmounts[cardToValue(c)]++
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cardAmounts)))
	cardAmounts[0] += jokers

	if cardAmounts[0] == 5 { // five of a kind
		return 6
	}
	if cardAmounts[0] == 4 { // four of a kind
		return 5
	}
	if cardAmounts[0] == 3 && cardAmounts[1] == 2 { // full house
		return 4
	}
	if cardAmounts[0] == 3 { // three of a kind
		return 3
	}
	if cardAmounts[0] == 2 && cardAmounts[1] == 2 { // two pairs
		return 2
	}
	if cardAmounts[0] == 2 { // one pair
		return 1
	}
	return 0 // High Cards
}

func cardToValue(c rune) int {
	if c <= '9' && c >= '0' {
		return int(c - '0')
	}

	switch c {
	case 'T':
		return 10
	case 'J':
		return -1
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	}

	panic(fmt.Errorf("Unknown value %s", string(c)))
}
