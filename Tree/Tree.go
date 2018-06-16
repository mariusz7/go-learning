package main

import (
	"errors"
	"fmt"
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
		} else {
			if nil == t.head.right {
				t.head = t.head.left
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
		}
	} else {
		deleteImpl(value, t.head)
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
				} else {
					if nil == deletedNode.right {
						curRoot.left = deletedNode.left
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

				if deletedNode.left == nil {
					curRoot.right = deletedNode.right
				} else {
					if deletedNode.right == nil {

						curRoot.right = deletedNode.left
					} else {

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
	}

	if t.head == nil {
		return 0, errors.New("Tree does not contain any elements")
	}

	var curElement int = 0

	return getNthElementImpl(n, & curElement, t.head)
}

func getNthElementImpl(n int, curElement *int, node *Node) (int, error) {

	if node.left != nil {
		if v, e := getNthElementImpl(n, curElement, node.left); e == nil {
			return v, nil
		}
	}

	if n == *curElement {
		return node.value, nil
	}

	*curElement++

	if node.right != nil {
		if v, e := getNthElementImpl(n, curElement, node.right); e == nil {
			return v, nil
		}
	}

	return 0, errors.New("value requested not present in this tree")
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

		//fmt.Printf("allValues after appending: %v\n", allValues)

		if nil != node.right {
			//3. Append values from the right subtree
			allValues = getAllValues(node.right, allValues)
		}
	}

	return allValues
}
