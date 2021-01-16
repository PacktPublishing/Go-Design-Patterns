package main

import "fmt"


func main(){
	addN := func(m int)func(int)int{
		return func(n int)int{
			return m+n
		}
	}
	addFive := addN(5)
	result := addFive(6)
	fmt.Println(result)
}