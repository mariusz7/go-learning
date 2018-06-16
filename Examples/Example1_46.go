package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	sum, product := foo(array[:])
	fmt.Printf ("%v, %v", sum, product)
}

func foo(array []int) (int, int) {

	var sum, product int = 0, 1
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}

	for i := 0; i < len(array); i++ {
		product *= array[i]
	}

	return sum, product
}
