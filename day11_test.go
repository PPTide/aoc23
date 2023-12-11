package main

import "testing"

func TestMainDay11(t *testing.T) {
	in := "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#....."

	out := mainDay11(in)
	expectedOut := "374"

	if out != expectedOut {
		t.Fatalf("expected output \"%s\" got \"%s\"", expectedOut, out)
	}
}
