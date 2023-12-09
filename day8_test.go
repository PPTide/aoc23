package main

import "testing"

func TestMainDay8(t *testing.T) {
	in := "LR\n\n11A = (11B, XXX)\n11B = (XXX, 11Z)\n11Z = (11B, XXX)\n22A = (22B, XXX)\n22B = (22C, 22C)\n22C = (22Z, 22Z)\n22Z = (22B, 22B)\nXXX = (XXX, XXX)"

	out := mainDay8(in)
	expectedOut := "Please find the LCM: [2 3]"

	if out != expectedOut {
		t.Fatalf("expected output \"%s\" got \"%s\"", expectedOut, out)
	}
}
