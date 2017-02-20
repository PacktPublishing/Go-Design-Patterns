package main

import (
	"fmt"
	"reflect"
)

func main() {
	//Explicitly declaring a "string" variable
	var explicit string = "Hello, I'm a explicitly declared variable"

	//Implicitly declaring a "string"
	inferred := "Hello World!"

	fmt.Println("Variable 'explicit' is of type:", reflect.TypeOf(explicit))
    fmt.Println("Variable 'inferred' is of type:", reflect.TypeOf(inferred))

}
