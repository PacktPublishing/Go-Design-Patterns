package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"runtime"
)

const size = 9999
const (
	xSize    = 2560
	ySize    = 1920
	channels = 4
)

type Image [channels][xSize][ySize]uint32

type bigMatrix [size][size]float64

var bigMatrixArray []bigMatrix
var matrix bigMatrix

func main() {
	fmt.Printf("Enter an option between \"add\" and \"total\". Use \"quit\" to exit\n")

	for {
		input := new(string)
		fmt.Fscanf(os.Stdin, "%s\n", input)
		switch *input {
		case "add":
			createNew()
		case "total":
			fmt.Printf("%d objects created\n", len(bigMatrixArray))
		case "quit":
			fmt.Println("Quitting")
			os.Exit(0)
		default:
			fmt.Printf("%d objects created\n", len(bigMatrixArray))
		}
	}
}

func createNew() {
	var mem runtime.MemStats

	runtime.ReadMemStats(&mem)
	fmt.Printf("Alloc: %d MB, ", mem.Alloc/1024/1024)
	fmt.Printf("Total Alloc: %d MB, ", mem.TotalAlloc/1024/1024)
	fmt.Printf("HeapAlloc: %d MB, ", mem.HeapAlloc/1024/1024)
	fmt.Printf("HeapSys: %d MB\n", mem.HeapSys/1024/1024)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i][j] = float64(i) + float64(j)
		}
	}

	bigMatrixArray = append(bigMatrixArray, matrix)

	runtime.ReadMemStats(&mem)
	fmt.Printf("Alloc: %d MB, ", mem.Alloc/1024/1024)
	fmt.Printf("Total Alloc: %d MB, ", mem.TotalAlloc/1024/1024)
	fmt.Printf("HeapAlloc: %d MB, ", mem.HeapAlloc/1024/1024)
	fmt.Printf("HeapSys: %d MB\n\n", mem.HeapSys/1024/1024)
}
