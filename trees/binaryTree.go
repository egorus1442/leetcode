package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// вставка в бинарное дерево
func (n *Node) Insert(value int) {
	if n == nil {
		return
	}

	if value < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

// поиск по значению в бинарном дереве
func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}

	if value == n.Value {
		return true
	} else if value < n.Value {
		return n.Left.Search(value)
	} else {
		return n.Right.Search(value)
	}
}

// прямой обход бинарного дерева
func (n *Node) PreOrderTraversal() error {
	if n == nil {
		return errors.New("not foun a Node")
	}

	fmt.Printf("%d", n.Value)
	n.Left.PreOrderTraversal()
	n.Right.PreOrderTraversal()
	return nil
}

func main() {
	root := &Node{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)

	fmt.Print(root.Search(4))
}
