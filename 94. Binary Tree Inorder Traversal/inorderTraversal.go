package main

//прямой рекурсивный обход

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	traverse(root, &result)
	return result
}

func traverse(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	traverse(node.Left, result)
	*result = append(*result, node.Val)
	traverse(node.Right, result)
}

func main() {
	// Пример дерева: [1, null, 2, 3]
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	resultRecursive := inorderTraversal(root)
	fmt.Println("Рекурсивный подход:", resultRecursive)
}
