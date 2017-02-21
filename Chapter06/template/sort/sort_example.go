package main

import (
	"sort"
	"fmt"
)

type MyList []int

func (m MyList) Len() int {
	return len(m)
}

func (m MyList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MyList) Less(i, j int) bool {
return m[i] < m[j]
}

func main() {
	var myList MyList = []int{6,4,2,8,1}

	fmt.Println(myList)
	sort.Sort(myList)
	fmt.Println(myList)
}
