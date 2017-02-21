package arithmetic

import "testing"

func TestSum(t *testing.T) {
	a := 5
	b := 6
	expected := 11

	res := Sum(a, b)
	if res != expected {
		t.Errorf("%d+%d isn't %d\n", a, b, res)
	}

	res = Sum()
	if res != 0 {
		t.Errorf("Calling Sum function without values must return 0, not %d", res)
	}
}

func TestSubtract(t *testing.T) {
	a := 3
	b := 8
	expected := -5

	res := Subtract(a, b)
	if res != expected {
		t.Errorf("%d-%d isn't %d\n", a, b, res)
	}

	expected = 0
	res = Subtract(a)
	if res != expected {
		t.Errorf("When passing only one argumment to subtract it must return 0")
	}
}

func TestMultiply(t *testing.T) {
	a := 8
	b := 7
	expected := 56

	res := Multiply(a, b)
	if res != expected {
		t.Errorf("%d-%d isn't %d\n", a, b, res)
	}

	expected = 0
	res = Multiply(a)
	if res != expected {
		t.Errorf("When passing only one argument to multiply it must return 0")
	}
}

func TestDivide(t *testing.T) {
	a := 10
	b := 2
	expected := 5.0

	//Test a common division
	res, err := Divide(a, b)
	if err != nil {
		t.Error(err)
	}
	if res != expected {
		t.Errorf("%d-%d isn't %f\n", a, b, res)
	}

	//Test a division by zero that must return an error
	b = 0
	res, err = Divide(a, b)
	if err == nil {
		t.Error("No error returned after passed 0 as divisor")
	}
}
