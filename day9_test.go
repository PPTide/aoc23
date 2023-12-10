package main

import "testing"

func TestMainDay9(t *testing.T) {
	in := "0 3 6 9 12 15\n1 3 6 10 15 21\n10 13 16 21 30 45"

	out := mainDay9(in)
	expectedOut := "2"

	if out != expectedOut {
		t.Fatalf("expected output \"%s\" got \"%s\"", expectedOut, out)
	}
}
