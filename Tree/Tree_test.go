package main

import "testing"

func compareEvsA(expected []int, actual []int, t *testing.T) {
	if len(expected) != len(actual) {
		t.Errorf("Expected len(%v): %v != actual len(%v): %v\n", expected, len(expected), actual, len(actual))
	} else {
		for i := 0; i < len(expected); i++ {
			if expected[i] != actual[i] {
				t.Errorf("expected %v != actual %v at position %v\n", expected, actual, i)
			}
		}
	}
}

func T1estGetAllValues(t *testing.T) {

/*
                2
               / \
              /   \
             /     \
            1       7
                   / \
                  /   \
                 /     \
                3       10
               / \     / \
              2   5   9   10
                           \
                            10
*/

	var testNodes []Node = []Node{
		Node{value: 2},  //0
		Node{value: 7},  //1
		Node{value: 10}, //2
		Node{value: 10}, //3
		Node{value: 1},  //4
		Node{value: 3},  //5
		Node{value: 5},  //6
		Node{value: 10}, //7
		Node{value: 9},  //8
		Node{value: 2},   //9
	}

	//2
	testNodes[0].left = &testNodes[4]
	testNodes[0].right = &testNodes[1]

	//7
	testNodes[1].left = &testNodes[5]
	testNodes[1].right = &testNodes[2]

	//10
	testNodes[2].left = &testNodes[8]
	testNodes[2].right = &testNodes[3]

	//10
	testNodes[3].left = nil
	testNodes[3].right = &testNodes[7]

	//1
	testNodes[4].left = nil
	testNodes[4].right = nil

	//3
	testNodes[5].left = &testNodes[9]
	testNodes[5].right = &testNodes[6]

	//5
	testNodes[6].left = nil
	testNodes[6].right = nil

	//10
	testNodes[7].left = nil
	testNodes[7].right = nil

	//9
	testNodes[8].left = nil
	testNodes[8].right = nil

	//2
	testNodes[9].left = nil
	testNodes[9].right = nil

	var tree Tree
/*
	for _, node : range testNodes {
		tree.Insert(node.value)
	}
*/
	for i := 0; i < len(testNodes); i++ {
		tree.Insert(testNodes[i].value)
	}

	var expected []int = []int{1, 2, 2, 3, 5, 7, 9, 10, 10, 10}
	var actual []int = tree.GetAllValues()

	compareEvsA(expected, actual, t)
}

func T1estInsert(t *testing.T) {

	var inserted []int = []int{3, 1, 4, 8, 5, 1, 5, 2, 1, 8}
	var expected []int = []int{1, 1, 1, 2, 3, 4, 5, 5, 8, 8}

	var tree Tree

/*
	for _, v : range inserted {
		tree.Insert(v)
	}
*/
	for i := 0; i < len(inserted); i++ {
		tree.Insert(inserted[i])
	}

	compareEvsA(expected, tree.GetAllValues(), t)
}

func TestDelete(t *testing.T) {

	var inserted []int = []int{3, 1, 4, 8, 5, 1, 5, 2, 1, 8}
	var expected []int = []int{1, 1, 1, 2, 3, 4, 5, 5, 8, 8}

	var tree Tree
/*
	for _, v : range inserted {
		tree.Insert(v)
	}
*/
	for i := 0; i < len(inserted); i++ {
		tree.Insert(inserted[i])
	}

	tree.PrintTree()

	compareEvsA(expected, tree.GetAllValues(), t)

	tree.Delete(1)
	expected = []int{1, 1, 2, 3, 4, 5, 5, 8, 8}
	compareEvsA(expected, tree.GetAllValues(), t)

	tree.Delete(5)
	expected = []int{1, 1, 2, 3, 4, 5, 8, 8}
	compareEvsA(expected, tree.GetAllValues(), t)

	tree.Delete(8)
	expected = []int{1, 1, 2, 3, 4, 5, 8}
	compareEvsA(expected, tree.GetAllValues(), t)

	tree.Delete(8)
	expected = []int{1, 1, 2, 3, 4, 5}
	compareEvsA(expected, tree.GetAllValues(), t)
/*
	tree.Delete(1)
	expected = []int{1, 2, 3, 4, 5}
	compareEvsA(expected, tree.GetAllValues(), t)

	tree.Delete(1)
	expected = []int{2, 3, 4, 5}
	compareEvsA(expected, tree.GetAllValues(), t)

	tree.Delete(4)
	expected = []int{2, 3, 5}
	compareEvsA(expected, tree.GetAllValues(), t)

	tree.Delete(5)
	expected = []int{2, 3}
	compareEvsA(expected, tree.GetAllValues(), t)

	tree.Delete(3)
	expected = []int{2}
	compareEvsA(expected, tree.GetAllValues(), t)

	tree.Delete(2)
	expected = []int{}
	compareEvsA(expected, tree.GetAllValues(), t)
*/
}
/*
func TestGetNthElement(t *testing.T) {

	var inserted []int = []int{3, 1, 4, 8, 5, 1, 5, 2, 1, 8}
	var expected []int = []int{1, 1, 1, 2, 3, 4, 5, 5, 8, 8}

	var tree Tree


//	for i, v : range inserted {
//		tree.Insert(v)
//
//		if vreturned, err := tree.GetNthElement(i); err != nil || vreturned != v {
//			t.Errorf("Failed to retrieve just inserted value %v at position %v\n", v, i)
//		}
//	}

	for i := 0; i < len(inserted); i++ {
		tree.Insert(inserted[i])

		if vreturned, err := tree.GetNthElement(i); err != nil || vreturned != inserted[i] {
			t.Errorf("Failed to retrieve just inserted value %v at position %v\n", inserted[i], i)
		}
	}

//	for i, v : range expected {
//		if vreturned, err := tree.GetNthElement(i); err != nil || vreturned != expected[i] {
//			t.Errorf("Failed to retrieve expected value %v at position %v in actual table %v\n", v, i, tree.GetAllValues())
//		}
//	}

	for i := 0; i < len(expected); i++ {
		if vreturned, err := tree.GetNthElement(i); err != nil || vreturned != expected[i] {
			t.Errorf("Failed to retrieve expected value %v at position %v in actual table %v\n", expected[i], i, tree.GetAllValues())
		}
	}
}
*/