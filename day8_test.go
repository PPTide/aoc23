package main

import "testing"

func TestMainDay8(t *testing.T) {
	in := "LLR\n\nAAA = (BBB, BBB)\nBBB = (AAA, ZZZ)\nZZZ = (ZZZ, ZZZ)"

	out := mainDay8(in)
	expectedOut := "8"

	if out != expectedOut {
		t.Fatalf("expected output \"%s\" got \"%s\"", expectedOut, out)
	}
}
