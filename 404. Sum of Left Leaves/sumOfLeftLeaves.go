package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	answer := searchLefts(root)
	var sum int = 0
	for _, i := range answer {
		sum += i
	}
	return sum
}

func searchLefts(root *TreeNode) []int {
	var answer []int
	if root == nil {
		return answer
	}
	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			answer = append(answer, root.Left.Val)
		}
	}

	answer = append(answer, searchLefts(root.Right)...)
	answer = append(answer, searchLefts(root.Left)...)

	return answer
}
