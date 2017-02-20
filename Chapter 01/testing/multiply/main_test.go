package main

import "testing"

func TestMultiply(t *testing.T) {
	a := 5
	b := 6
	expected := 30

	res := Multiply(a, b)
	if res != expected {
		t.Errorf("Our multiply function doens't work, %d+%d isn't %d\n", a, b, res)
	}
}
