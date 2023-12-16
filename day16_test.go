package main

import "testing"

func TestMainDay16(t *testing.T) {
	in := ".|...\\....\n|.-.\\.....\n.....|-...\n........|.\n..........\n.........\\\n..../.\\\\..\n.-.-/..|..\n.|....-|.\\\n..//.|...."

	out := mainDay16(in)
	expectedOut := "46"

	if out != expectedOut {
		t.Fatalf("expected output \"%s\" got \"%s\"", expectedOut, out)
	}
}
