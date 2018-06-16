package main

import (
	"errors"
	"fmt"
	"container/list"
)

type ITree interface {
	Insert(value int)
	Delete(value int)
	GetNthElement(n int) int
	GetAllValues() []int
}

type Tree struct {
	head *Node
}

type Node struct {
	left *Node
	right *Node
	value int
}

func newNode(v int) *Node {
	var node *Node = &Node{value: v}
	return node
}

func (t *Tree) Insert(value int) {
	if t.head == nil {
		t.head = newNode(value)
		//fmt.Printf("value in head: %v\n", t.head.value)
	} else {
		t.head.insertRec(value)
	}
}

func (n *Node) insertRec(value int) {

	//Duplicates will be inserted on the right side
	if value < n.value { //left
		if nil == n.left {
			n.left = newNode(value)
		} else {
			n.left.insertRec(value)
		}
	} else { //right
		if nil == n.right {
			n.right = newNode(value)
		} else {
			n.right.insertRec(value)
		}
	}
}

func assert(everythingOK bool) {
	if !everythingOK {
		panic("Assertion failed.")
	}
}

func (t *Tree) PrintTree() {
	var tree2D [][]string = [][]string{
		[]string{"1"},
		[]string{"2", "3"},
		[]string{"4", "5", "6", "7"},
	}

	var fifo *list.List = list.New()

	//Traverse the tree using breadth-first way
	if nil != t.head {
		fifo.PushBack(t.head)
		while fifo.Len() > 0 {
			var node *Node = fifo.Front()
		}
	}


	//Get 2D tree

	for row := 0; row < len(tree2D); row++ {
		fmt.Printf("row: %v\n", tree2D[row])
	}
}

func (t *Tree) printTree() {

		

}

// Deletes the leftmost value first, if there are duplicates
func (t *Tree) Delete(value int) {

	fmt.Printf("Requested to delete value: %v\n", value)


	if nil == t.head {
		return
	} else if t.head.value == value {
		var deletedNode = t.head
		fmt.Printf("deletedNode.value: %v\n", deletedNode.value)

		if nil == t.head.left { //There is no left subtree
			t.head = t.head.right
		} else if nil == t.head.right { //There are none subtrees
			//Nothing more to do, if this was the last node in the tree
		} else { //Both subtrees exist
			//2 options available:
			// a) make left subtree part of the right subtree
			// b) make right subtree part of the left subtree

			//Let's implement 2.a
			t.head = t.head.right

			//In the right subtree find the leftmost node with 'left' field equal to nil
			var lnode *Node = deletedNode.right
			for ; lnode.left != nil; lnode = lnode.left {
			}

			lnode.left = deletedNode.left

			//Symbolic free. Fairwell deletedNode...
			deletedNode = nil
		}
	} else {
		var curRoot *Node = t.head
		deleteImpl(value, curRoot)
	}
}

func deleteImpl(value int, curRoot *Node) {

	assert(curRoot.value != value)

	if value < curRoot.value { //Need to delete from the left subtree

		if nil == curRoot.left {
			//That value is not present in the tree
		} else {

			if value == curRoot.left.value {
				var deletedNode *Node = curRoot.left
				fmt.Printf("deletedNode.value: %v\n", deletedNode.value)

				if nil == deletedNode.left {
					curRoot.left = deletedNode.right
				} else if nil == deletedNode.right {
					//deletedNode does not have any subtrees we need to take care of
				} else { //Both subtrees exist

					curRoot.left = deletedNode.right

					//In the right subtree find the leftmost node with 'left' field equal to nil
					var lnode *Node = deletedNode.right
					deletedNode.right = nil //Just for sanity check later on

					for ; lnode.left != nil; lnode = lnode.left {
					}

					lnode.left = deletedNode.left
					deletedNode.left = nil //Just for sanity check later on

					//Symbolic free. Fairwell deletedNode...
					assert(nil == deletedNode.left && nil == deletedNode.right)
					deletedNode = nil
				}
			} else {
				deleteImpl(value, curRoot.left)
			}
		}
	} else { //Need to delete from the right subtree

		if nil == curRoot.right {
			//That value is not present in the tree
			fmt.Printf("That value is not present in the tree: %v\n", value)
		} else {

			if value == curRoot.right.value {
				var deletedNode *Node = curRoot.right
				fmt.Printf("deletedNode.value: %v\n", deletedNode.value)

				if nil == deletedNode.left {
					curRoot.right = deletedNode.right
				} else if nil == deletedNode.right {
					//deletedNode does not have any subtrees we need to take care of
				} else { //Both subtrees exist

					curRoot.right = deletedNode.right

					//In the right subtree find the leftmost node with 'left' field equal to nil
					var lnode *Node = deletedNode.right
					deletedNode.right = nil //Just for sanity check later on

					for ; lnode.left != nil; lnode = lnode.left {
					}

					lnode.left = deletedNode.left
					deletedNode.left = nil //Just for sanity check later on

					//Symbolic free. Fairwell deletedNode...
					assert(nil == deletedNode.left && nil == deletedNode.right)
					deletedNode = nil
				}
			} else {
				deleteImpl(value, curRoot.right)
			}
		}
	}
}

func (t *Tree) GetNthElement(n int) (int, error) {

	if n < 0 {
		return 0, errors.New("n < 0")
	} else if t.head == nil {
		return 0, errors.New("Tree does not contain any elements")
	}

	var curElement int //holds 0

	return getNthElementImpl(n, & curElement, t.head)
}

func getNthElementImpl(n int, curElement *int, node *Node) (int, error) {

	if node.left == nil {
		if n == *curElement {
			return node.value, nil
		} else {
			if node.right == nil {
				return 0, errors.New("Dead end")
			} else {
				*curElement++
				return getNthElementImpl(n, curElement, node.right)
			}
		}
	} else {
		if v, err := getNthElementImpl(n, curElement, node.left); err == nil {
			return v, err
		} else {
			if n == *curElement {
				return node.value, nil
			} else {
				*curElement++
				return getNthElementImpl(n, curElement, node.right)
			}
		}
	}
}

func (t *Tree) GetAllValues() []int {

	return getAllValues(t.head, []int{})
}

func getAllValues(node *Node, allValues []int) []int {

	if node != nil {
		if nil != node.left {
			//1. Append values from the left subtree
			allValues = getAllValues(node.left, allValues)
		}

		//2. Append the current value
		allValues = append(allValues, node.value)

		fmt.Printf("allValues after appending: %v\n", allValues)

		if nil != node.right {
			//3. Append values from the right subtree
			allValues = getAllValues(node.right, allValues)
		}
	}

	return allValues
}
