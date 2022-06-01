package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {

	head := root

	if head == nil {
		return head
	}

	// head variable is needed so that the left and right pointers of root are preserved
	tempRight := head.Right
	tempLeft := head.Left
	root.Left = invertTree(tempRight)
	root.Right = invertTree(tempLeft)

	return root
}
