package main

import "fmt"

func main(){
	add := func (m int)int{
		return m+1
	}
	result := add(6)
	// 1+6 must print 7
	fmt.Println(result)
}

