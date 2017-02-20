package arithmetic

import "errors"

// Sum will take as many arguments as you pass to it and return the sum of all
// of them
func Sum(args ...int) (res int) {
	for _, v := range args {
		res += v
	}

	return
}

// Subtract will take as many arguments as you pass to it and return the
// subtraction of them
func Subtract(args ...int) int {
	if len(args) < 2 {
		return 0
	}

	res := args[0]
	for i := 1; i < len(args); i++ {
		res -= args[i]
	}

	return res
}

// Multiply will take as many arguments as you want and will return the
// multiplication of them
func Multiply(args ...int) int {
	if len(args) < 2 {
		return 0
	}

	res := 1
	for i := 0; i < len(args); i++ {
		res *= args[i]
	}

	return res
}

// Divide will take two arguments and return its division as far as the second
// argument isn't zero that will return an error
func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("You cannot divide by zero")
	}

	return float64(a) / float64(b), nil
}
