package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Explicitly declaring a "string" variable
	var explicitString string = "Hello World!"

	//Implicitly declaring a "string"
	implicitString := "Hello World!"

	println(explicitString)
	println(implicitString)

	fmt.Println("Variable 'implicitString' is of type:", reflect.TypeOf(implicitString))
	fmt.Println("Variable 'explicitString' is of type:", reflect.TypeOf(explicitString))
}
