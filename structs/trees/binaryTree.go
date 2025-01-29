package trees

import (
	"fmt"
)

// Узел бинарного дерева
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Вставка нового значения в бинарное дерево
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

// Поиск значения в бинарном дереве
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

// Прямой обход (Pre-order traversal) // Корень-левый-правый
func (n *Node) PreOrderTraversal() {
	if n == nil {
		return
	}

	fmt.Printf("%d ", n.Value)
	n.Left.PreOrderTraversal()
	n.Right.PreOrderTraversal()
}

// Симметричный обход (In-order traversal)
func (n *Node) InOrderTraversal() {
	if n == nil {
		return
	}

	n.Left.InOrderTraversal()
	fmt.Printf("%d ", n.Value)
	n.Right.InOrderTraversal()
}

// Обратный обход (Post-order traversal) //левое поддерево -> корень -> правое поддерево
func (n *Node) PostOrderTraversal() {
	if n == nil {
		return
	}

	n.Left.PostOrderTraversal()
	n.Right.PostOrderTraversal()
	fmt.Printf("%d ", n.Value)
}

func main() {
	// Создаем корень бинарного дерева
	root := &Node{Value: 10}

	// Вставляем значения в дерево
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(12)
	root.Insert(18)

	// Поиск значения в дереве
	fmt.Println("Поиск значения 7:", root.Search(7))
	fmt.Println("Поиск значения 20:", root.Search(20))

	// Обход дерева
	fmt.Print("Прямой обход: ")
	root.PreOrderTraversal()
	fmt.Println()

	fmt.Print("Симметричный обход: ")
	root.InOrderTraversal()
	fmt.Println()

	fmt.Print("Обратный обход: ")
	root.PostOrderTraversal()
	fmt.Println()
}
