package main

import "fmt"

func main() {
	printPairs([]int{1, 2, 3, 4, 5})
}

func printPairs(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			fmt.Printf("%v, %v\n", array[i], array[j])
		}
	}
}
