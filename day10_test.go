package main

import "testing"

func TestMainDay10(t *testing.T) {
	in := "7-F7-\n.FJ|7\nSJLL7\n|F--J\nLJ.LJ"

	out := mainDay10(in)
	expectedOut := "8"

	if out != expectedOut {
		t.Fatalf("expected output \"%s\" got \"%s\"", expectedOut, out)
	}
}
