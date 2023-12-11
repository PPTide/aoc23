package main

import "testing"

func TestMainDay11(t *testing.T) {
	in := "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#....."

	out := mainDay11(in, 2)
	expectedOut := "374"

	if out != expectedOut {
		t.Fatalf("expected output \"%s\" got \"%s\"", expectedOut, out)
	}

	out = mainDay11(in, 10)
	expectedOut = "1030"

	if out != expectedOut {
		t.Fatalf("expansionVal 10: expected output \"%s\" got \"%s\"", expectedOut, out)
	}

	out = mainDay11(in, 100)
	expectedOut = "8410"

	if out != expectedOut {
		t.Fatalf("expansionVal 100: expected output \"%s\" got \"%s\"", expectedOut, out)
	}
}
