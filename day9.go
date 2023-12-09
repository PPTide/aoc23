package main

import (
	"strconv"
	"strings"
)

func mainDay9(in string) string {
	sum := 0

	for _, line := range strings.Split(in, "\n") {
		derivatives := make([][]int, 0)

		derivatives = append(derivatives, Map(strings.Split(line, " "), func(t string) int {
			i, err := strconv.Atoi(t)

			if err != nil {
				panic(err)
			}

			return i
		}))
		for {
			der, allZeros := derive(derivatives[len(derivatives)-1])
			derivatives = append(derivatives, der)
			if allZeros {
				break
			}
		}
		//fmt.Println(derivatives)
		prevVal := 0
		for i := len(derivatives) - 1; i >= 0; i-- {
			prevVal = derivatives[i][0] - prevVal
		}

		sum += prevVal
	}

	return strconv.Itoa(sum)
}

func derive(in []int) (out []int, allZeros bool) {
	out = make([]int, 0, len(in)-1)

	allZeros = true
	for i := 1; i < len(in); i++ {
		diff := in[i] - in[i-1]
		out = append(out, diff)
		if diff != 0 {
			allZeros = false
		}
	}

	return
}
