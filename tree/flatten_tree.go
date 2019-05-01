package tree

import "github.com/mngoutham/GOExamples"

func flatten(root *GOExamples.TreeNode) *GOExamples.TreeNode {
	if root == nil {
		return root
	}
	left := flatten(root.Left)
	right := flatten(root.Right)
	curr := root
	curr.Right = left
	curr.Left = nil
	for ; curr.Right != nil; curr = curr.Right {}
	curr.Right = right
	return root
}
