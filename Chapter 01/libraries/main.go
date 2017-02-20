package main

import (
	"fmt"

	"github.com/sayden/go-design-patterns/introduction/libraries/arithmetic"
)

func main() {
	sumRes := arithmetic.Sum(5, 6)
	subRes := arithmetic.Subtract(10, 5)
	multiplyRes := arithmetic.Multiply(8, 7)
	divideRes, err := arithmetic.Divide(10, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("5+6 is %d. 10-5 is %d, 8*7 is %d and 10/2 is %f\n", sumRes, subRes, multiplyRes, divideRes)
}
