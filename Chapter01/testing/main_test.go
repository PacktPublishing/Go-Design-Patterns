package main

import "testing"

func TestSum(t *testing.T) {
	a := 5
	b := 6
	expected := 11

	res := sum(a, b)
	if res != expected {
		t.Errorf("Our sum function doens't work, %d+%d isn't %d\n", a, b, res)
	}
}
