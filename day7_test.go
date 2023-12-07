package main

import "testing"

func TestMainDay7(t *testing.T) {
	in := "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483"

	out := mainDay7(in)
	expectedOut := "5905"

	if out != expectedOut {
		t.Fatalf("expected output \"%s\" got \"%s\"", expectedOut, out)
	}
}
