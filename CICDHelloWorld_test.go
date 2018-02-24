package main

import "testing"

func TestCICD(t *testing.T) {
	a, b, c := Content()
	x, y, z := Content()
	if x != a {
		t.Errorf("This is meant to pass (for now)")
	}
	if y != b {
		t.Errorf("This is meant to pass (for now)")
	}
	if c != z {
		t.Errorf("This is meant to pass (for now)")
	}
}
