package interpreter

import "testing"

func TestCalculate(t *testing.T) {
	tempOperation := "5 3 sub 8 mul 4 sum 5 div"
	res, err := Calculate(tempOperation)
	if err != nil {
		t.Error(err)
	}

	if res != 4 {
		t.Errorf("Expected result not found: %d != %d\n", 4, res)
	}

	tempOperation = "3 4 sum 2 sub"
	res, err = Calculate(tempOperation)
	if err != nil {
		t.Error(err)
	}

	if res != 5 {
		t.Errorf("Expected result not found: %d != %d\n", 5, res)
	}
}
