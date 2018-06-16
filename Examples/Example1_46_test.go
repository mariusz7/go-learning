package main

import "testing"

func TestFoo(t *testing.T) {

	array := []int{1, 2, 3, 4, 5}
	expectedSum := 15
	expectedProduct := 120
	if actualSum, actualProduct := foo(array); expectedSum != actualSum || expectedProduct != actualProduct {
		t.Errorf("Invalid results from foo(). Expected (%v, %v), but got (%v, %v) for array: %v\n",
			expectedSum, expectedProduct, actualSum, actualProduct, array)
	}

	array = []int{0, 0, 0, 0, 0}
	expectedSum = 0
	expectedProduct = 0
	if actualSum, actualProduct := foo(array); expectedSum != actualSum || expectedProduct != actualProduct {
		t.Errorf("Invalid results from foo(). Expected (%v, %v), but got (%v, %v) for array: %v\n",
			expectedSum, expectedProduct, actualSum, actualProduct, array)
	}

	array = []int{1, 1, 1, 1, 1}
	expectedSum = 5
	expectedProduct = 1
	if actualSum, actualProduct := foo(array); expectedSum != actualSum || expectedProduct != actualProduct {
		t.Errorf("Invalid results from foo(). Expected (%v, %v), but got (%v, %v) for array: %v\n",
			expectedSum, expectedProduct, actualSum, actualProduct, array)
	}

	array = []int{-1, 1, -1, 1, 0}
	expectedSum = 0
	expectedProduct = 0
	if actualSum, actualProduct := foo(array); expectedSum != actualSum || expectedProduct != actualProduct {
		t.Errorf("Invalid results from foo(). Expected (%v, %v), but got (%v, %v) for array: %v\n",
			expectedSum, expectedProduct, actualSum, actualProduct, array)
	}

	array = []int{-99999, 100000, 0, -1, 1}
	expectedSum = 1
	expectedProduct = 0
	if actualSum, actualProduct := foo(array); expectedSum != actualSum || expectedProduct != actualProduct {
		t.Errorf("Invalid results from foo(). Expected (%v, %v), but got (%v, %v) for array: %v\n",
			expectedSum, expectedProduct, actualSum, actualProduct, array)
	}
}
