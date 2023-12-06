package main

import (
	"testing"
)

func TestMainDay6(t *testing.T) {
	in := "Time:      7  15   30\nDistance:  9  40  200"

	out := mainDay6(in)

	if out != "288" {
		t.Fatalf("expected output \"288\" got \"%s\"", out)
	}
}
