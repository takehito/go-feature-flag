package main

import "testing"

func TestSum(t *testing.T) {
	a := 1
	b := 2
	expected := 3
	if got := sum(a, b); got != expected {
		t.Errorf("expected: %d, got: %d\n", expected, got)
	}
}
