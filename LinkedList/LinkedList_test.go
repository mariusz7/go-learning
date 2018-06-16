package main

import "testing"

func testArrays(expectedList, actualList []int, list *LinkedList, t *testing.T) {
	if len(expectedList) != len(actualList) {
		t.Errorf("Actual list (%v) has different length (%v) than expected (%v) for (%v)\n",
			actualList, len(actualList), len(expectedList), expectedList)
	} else {
		for i := 0; i < len(expectedList); i++ {
			if expectedList[i] != actualList[i] {
				t.Errorf("Actual list %v is different from the expected one %v\n",
					actualList, expectedList)
			}
		}
	}
}

func TestAddNodeEnd(t *testing.T) {
	var list LinkedList
	for i := 10; i < 20; i++ {
		list.AddNodeEnd(i)
	}

	var expectedList []int = []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	var actualList []int = list.GetAllNodes()

	testArrays(expectedList, actualList, &list, t)
}

func TestAddNodeBegin(t *testing.T) {
	var list LinkedList
	for i := 10; i < 20; i++ {
		list.AddNodeBegin(i)
	}

	var expectedList []int = []int{19, 18, 17, 16, 15, 14, 13, 12, 11, 10}
	var actualList []int = list.GetAllNodes()

	testArrays(expectedList, actualList, &list, t)
}

func TestAddNodeAfterNode(t *testing.T) {
	var list LinkedList
	for i := 10; i < 20; i++ {
		list.AddNodeEnd(i)
	}

	var n *Node = list.GetNthNode(5)
	list.AddNodeAfterNode(n, 100)

	n = list.GetNthNode(10)
	list.AddNodeAfterNode(n, 200)
	
	var expectedList []int = []int{10, 11, 12, 13, 14, 15, 100, 16, 17, 18, 19, 200}
	var actualList []int = list.GetAllNodes()

	testArrays(expectedList, actualList, &list, t)
}

func TestAddNodeBeforeNode(t *testing.T) {
	var list LinkedList
	for i := 10; i < 20; i++ {
		list.AddNodeEnd(i)
	}

	var n *Node = list.GetNthNode(0)
	list.AddNodeBeforeNode(n, 300)

	n = list.GetNthNode(0)
	list.AddNodeBeforeNode(n, 400)

	n = list.GetNthNode(11)
	list.AddNodeBeforeNode(n, 500)
	
	var expectedList []int = []int{400, 300, 10, 11, 12, 13, 14, 15, 16, 17, 18, 500, 19}
	var actualList []int = list.GetAllNodes()

	testArrays(expectedList, actualList, &list, t)
}
