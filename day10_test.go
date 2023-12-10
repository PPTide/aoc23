package main

import "testing"

func TestMainDay10(t *testing.T) {
	in := "....\n.F-7\n.|.|\n.L7S\n.FJ|\n.L7|\n.FJ|\n.L7|\nF-J|\nL-7|\nF-J|\nL7.|\nFJ.|\nL7.|\n.L-J"
	//in := "7-F7-\n.FJ|7\nSJLL7\n|F--J\nLJ.LJ"
	//in := ".....\n.S-7.\n.|.|.\n.L-J.\n....."

	out := mainDay10(in)
	//expectedOut := "8"
	expectedOut := "4"

	if out != expectedOut {
		t.Fatalf("expected output \"%s\" got \"%s\"", expectedOut, out)
	}
}
