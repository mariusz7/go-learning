package main

import (
	"testing"
	"fmt"
)

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

func TestGetAllValues(t *testing.T) {

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
		Node{value: 2},  //9
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

	for _, node := range testNodes {
		tree.Insert(node.value)
	}

	var expected []int = []int{1, 2, 2, 3, 5, 7, 9, 10, 10, 10}
	var actual []int = tree.GetAllValues()

	compareEvsA(expected, actual, t)
}

func TestInsert(t *testing.T) {

	var inserted []int = []int{3, 1, 4, 8, 5, 1, 5, 2, 1, 8}
	var expected []int = []int{1, 1, 1, 2, 3, 4, 5, 5, 8, 8}

	var tree Tree

	for _, v := range inserted {
		tree.Insert(v)
	}

	compareEvsA(expected, tree.GetAllValues(), t)
}

func TestDelete(t *testing.T) {

	var inserted []int = []int{3, 1, 4, 8, 5, 1, 5, 2, 1, 8}
	var expected []int = []int{1, 1, 1, 2, 3, 4, 5, 5, 8, 8}

	var tree Tree
	compareEvsA([]int{}, tree.GetAllValues(), t)

	tree.Delete(666)
	//printTree(&tree)
	compareEvsA([]int{}, tree.GetAllValues(), t)

	for _, v := range inserted {
		tree.Insert(v)
	}

	//printTree(&tree)

	compareEvsA(expected, tree.GetAllValues(), t)

	tree.Delete(1)
	//printTree(&tree)
	expected = []int{1, 1, 2, 3, 4, 5, 5, 8, 8}
	compareEvsA(expected, tree.GetAllValues(), t)

	tree.Delete(5)
	//printTree(&tree)
	expected = []int{1, 1, 2, 3, 4, 5, 8, 8}
	compareEvsA(expected, tree.GetAllValues(), t)

	tree.Delete(8)
	//printTree(&tree)
	expected = []int{1, 1, 2, 3, 4, 5, 8}
	compareEvsA(expected, tree.GetAllValues(), t)

	tree.Delete(8)
	//printTree(&tree)
	expected = []int{1, 1, 2, 3, 4, 5}
	compareEvsA(expected, tree.GetAllValues(), t)

	tree.Delete(1)
	//printTree(&tree)
	expected = []int{1, 2, 3, 4, 5}
	compareEvsA(expected, tree.GetAllValues(), t)

	tree.Delete(1)
	//printTree(&tree)
	expected = []int{2, 3, 4, 5}
	compareEvsA(expected, tree.GetAllValues(), t)

	tree.Delete(1) //Trying to delete non-existing value on the left hand side

	tree.Delete(4)
	//printTree(&tree)
	expected = []int{2, 3, 5}
	compareEvsA(expected, tree.GetAllValues(), t)

	tree.Delete(5)
	//printTree(&tree)
	expected = []int{2, 3}
	compareEvsA(expected, tree.GetAllValues(), t)

	tree.Delete(3)
	//printTree(&tree)
	expected = []int{2}
	compareEvsA(expected, tree.GetAllValues(), t)

	tree.Delete(100) //Trying to delete non-existing value on the right hand side

	tree.Delete(2)
	//printTree(&tree)
	expected = []int{}
	compareEvsA(expected, tree.GetAllValues(), t)







	var tree2 Tree
	for _, v := range inserted {
		tree2.Insert(v)
	}
	//Missing case: delete root node having both subtrees
	tree2.Delete(3)
	expected = []int{1, 1, 1, 2, 4, 5, 5, 8, 8}
	compareEvsA(expected, tree2.GetAllValues(), t)


	var tree4 Tree
	for _, v := range []int{5, 3, 2, 1} {
		tree4.Insert(v)
	}
	//printTree(&tree4)
	//Missing case: delete node without right subtree
	tree4.Delete(2)
	expected = []int{1, 3, 5}
	compareEvsA(expected, tree4.GetAllValues(), t)


	var tree5 Tree
	for _, v := range []int{5, 3, 2, 1, 4} {
		tree5.Insert(v)
	}
	//printTree(&tree5)
	tree5.Delete(3)
	expected = []int{1, 2, 4, 5}
	compareEvsA(expected, tree5.GetAllValues(), t)





	var tree6 Tree
	for _, v := range []int{5, 3, 2, 1, 4, 10, 9, 8} {
		tree6.Insert(v)
	}
	//printTree(&tree6)
	tree6.Delete(5)
	expected = []int{1, 2, 3, 4, 8, 9, 10}
	compareEvsA(expected, tree6.GetAllValues(), t)





	var tree7 Tree
	for _, v := range []int{10, 5, 7, 4, 8, 9, 6} {
		tree7.Insert(v)
	}
	//printTree(&tree7)
	tree7.Delete(5)
	//printTree(&tree7)
	expected = []int{4, 6, 7, 8, 9, 10}
	compareEvsA(expected, tree7.GetAllValues(), t)




	var tree8 Tree
	for _, v := range []int{1, 10, 5, 15, 12} {
		tree8.Insert(v)
	}
	//printTree(&tree8)
	tree8.Delete(10)
	//printTree(&tree8)
	expected = []int{1, 5, 12, 15}
	compareEvsA(expected, tree8.GetAllValues(), t)
}

func TestGetNthElement(t *testing.T) {

	var inserted []int = []int{3, 1, 4, 8, 5, 1, 5, 2, 1, 8}
	var expected []int = []int{1, 1, 1, 2, 3, 4, 5, 5, 8, 8}

	var insertedInMeantime []int = []int{3, 3, 4, 8, 8, 8, 8, 8, 8, 8}

	var tree Tree
	if _, err := tree.GetNthElement(123); err == nil {
		t.Errorf("Getting from empty tree shall return an error\n")
	}

	if _, err := tree.GetNthElement(-123); err == nil {
		t.Errorf("Giving negative number as a parameter shall return an error\n")
	}

	for i, v := range inserted {
		tree.Insert(v)

		//printTree(&tree)

		vreturned, err := tree.GetNthElement(i)
		if err != nil {
			t.Errorf("err: %v\n", err)
		} else {
			if vreturned != insertedInMeantime[i] {
				t.Errorf("vreturned: %v != insertedInMeantime[%v]: %v\n",
					vreturned, i, v)
			}
		}
	}

	for i, v := range expected {
		if vreturned, err := tree.GetNthElement(i); err != nil || vreturned != expected[i] {
			t.Errorf("Failed to retrieve expected value %v at position %v in actual table %v\n", v, i, tree.GetAllValues())
		}
	}
}

func printTree(t *Tree) {

	//Holding pointers allows us to use 'nil' as a special value
	//for later formatting, meaning that that node is not present
	var tree2D [][]*int = [][]*int{}

	var curLevelNodes []*Node
	var nextLevelNodes []*Node

	//Traverse the tree using breadth-first way

	var visitNextLevelNodes bool //false by default

	if t.head != nil {
		nextLevelNodes = append(nextLevelNodes, t.head)
		visitNextLevelNodes = true
	}

	var level int = 0

	for visitNextLevelNodes {

		curLevelNodes = nextLevelNodes
		nextLevelNodes = nil
		visitNextLevelNodes = false
		tree2D = append(tree2D, []*int{})

		for _, node := range curLevelNodes {

			if node == nil {
				nextLevelNodes = append(nextLevelNodes, nil, nil)
				tree2D[level] = append(tree2D[level], nil)
			} else {
				nextLevelNodes = append(nextLevelNodes, node.left)
				nextLevelNodes = append(nextLevelNodes, node.right)
				tree2D[level] = append(tree2D[level], &node.value)

				if node.left != nil || node.right != nil {
					//There will be something to do for the next outer loop
					visitNextLevelNodes = true
				}
			}
		}

		level++
	}

	//Get 2D tree

	if len(tree2D) > 0 {

		const charsPerValue = 2
		const minSpaceBetweenValues = 2

		for row := 0; row < len(tree2D); row++ {

			var inBetweenSpacesPerLevel []int = []int{0, 30, 14, 6, 2}

			if row >= len(inBetweenSpacesPerLevel) {
				//Since this is only helper code to diagnoze bug in test, leave it for now
				//with check only - this is not production ready code yet
				panic("TODO implement supporting more levels")
			}

			//TODO fine tune
			var inBetweenSpaces int = inBetweenSpacesPerLevel[row]

			printLeftSpaces(row)

			for col := 0; col < len(tree2D[row]); col++ {
				if tree2D[row][col] != nil {
					printValue(*tree2D[row][col], charsPerValue)
				} else {
					printSpaces(charsPerValue)
				}

				//Don't print unneeded suffix after the last value in a row
				if col < len(tree2D[row]) - 1 {
					printSpaces(inBetweenSpaces)
				}
			}
			fmt.Println()

			printLines(row)
		}
	}
}

func printSpaces(n int) {
	//TODO find a better way of printing desired number of spaces...
	for i := 0; i < n; i++ {
		fmt.Print(" ")
	}
}

func printLeftSpaces(row int) {

	var leftSpaces []string = []string{
		"                              ",
		"              ",
		"      ",
		"  ",
		"",
	}

	if row >= len(leftSpaces) {
		panic("TODO implement support for more rows");
	} else {
		fmt.Print(leftSpaces[row])
	}
}

func printLines(row int) {
	var lines [][]string = [][]string{
		[]string{
			"                             /   \\",
			"                            /     \\",
			"                           /       \\",
			"                          /         \\",
			"                         /           \\",
			"                        /             \\",
			"                       /               \\",
			"                      /                 \\",
			"                     /                   \\",
			"                    /                     \\",
			"                   /                       \\",
			"                  /                         \\",
			"                 /                           \\",
			"                /                             \\",
		},
		[]string{
			"              / \\                             / \\",
			"             /   \\                           /   \\",
			"            /     \\                         /     \\",
			"           /       \\                       /       \\",
			"          /         \\                     /         \\",
			"         /           \\                   /           \\",
			"        /             \\                 /             \\",
		},
		[]string{
			"      / \\             / \\             / \\             / \\",
			"     /   \\           /   \\           /   \\           /   \\",
			"    /     \\         /     \\         /     \\         /     \\",
		},
		[]string{
			"  / \\     / \\     / \\     / \\     / \\     / \\     / \\     / \\",
		},
	}

	if row < len(lines) {
		for j := 0; j < len(lines[row]); j++ {
			fmt.Println(lines[row][j])
		}
	}
}

func printValue(v int, charsPerValue int) {

	if charsPerValue > 2 || v > 99 {
		panic("TODO add support for values grater than 99")
	}

	if v < 10 {
		fmt.Print(" ")
	}
	fmt.Print(v)
}
