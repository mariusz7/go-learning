package main

type LinkedList struct {
	head *Node
}

type Node struct {
	next *Node
	value int
}

func newNode(value int) *Node {
	var n *Node = new (Node)
	n.value = value

	return n
}

func (t *LinkedList) AddNodeEnd(value int) {
	if t.head == nil {
		t.head = newNode(value)
	} else {
		var n *Node = t.head
		for ; n.next != nil; n = n.next {
		}
		n.next = newNode(value)
	}
}

func (t *LinkedList) AddNodeBegin(value int) {

	var oldHead *Node = t.head
	t.head = newNode(value)
	t.head.next = oldHead
}

func (t *LinkedList) AddNodeAfterNode(n *Node, value int) {
	if n != nil {
		var oldNext *Node = n.next
		n.next = newNode(value)
		n.next.next = oldNext
	}
}

func (t *LinkedList) AddNodeBeforeNode(n *Node, value int) {
	if n != nil && t.head != nil {
		if n == t.head {
			var nn *Node = newNode(value)
			nn.next = t.head
			t.head = nn
		} else {
			var pn *Node
			for pn = t.head; pn != nil && pn.next != n; pn = pn.next {
			}

			if n == pn.next {
				var tmp *Node = newNode(value)
				pn.next = tmp
				tmp.next = n
			}
		}
	}
}

func (t *LinkedList) GetNthNode(index int) *Node {

	if t.head == nil {
		return nil
	}

	var n *Node = t.head
	var i int = 0

	for ; i < index; i++ {
		if n != nil {
			n = n.next
		} else {
			return nil
		}
	}

	return n
}

func (t *LinkedList) GetAllNodes() []int {
	var allNodes []int

	for n := t.head; n != nil; n = n.next {
		allNodes = append(allNodes, n.value)
	}
	
	return allNodes
}
