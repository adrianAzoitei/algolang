package main

import (
	"algorithms/karatsuba"
	"algorithms/mergesort"
	"fmt"
)

func main() {
	fmt.Println(karatsuba.Karatsuba(1234, 5678))
	fmt.Println(mergesort.MergeSort([]int{1, 3, 6, 7, 2, 4, 8, 5}))
}
