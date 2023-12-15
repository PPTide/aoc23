package main

import "testing"

func TestMainDay15(t *testing.T) {
	in := "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"

	out := mainDay15(in)
	expectedOut := "1320"

	if out != expectedOut {
		t.Fatalf("expected output \"%s\" got \"%s\"", expectedOut, out)
	}
}
