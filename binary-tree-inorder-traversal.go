package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var rec = func(root *TreeNode) {}
	rec = func(item *TreeNode) {
		if item == nil {
			return
		}

		fmt.Print(item.Val)

		if item.Left != nil {
			rec(item.Left)
		}

		if item.Right != nil {
			rec(item.Right)
		}

	}
	rec(root)
	return nil
}

func getTestData() *TreeNode {
	node3 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 2, Left: node3}
	root := &TreeNode{Val: 1, Right: node2}
	return root
}
