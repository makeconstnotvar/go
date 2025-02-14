package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Println(node.Val)

	if node.Left != nil {
		dfs(node.Left)
	}

	if node.Right != nil {
		dfs(node.Right)
	}
}

func inorderTraversal(root *TreeNode) []int {
	dfs(root)
	return nil
}

func getTestData1() *TreeNode {
	node3 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 2, Left: node3}
	root := &TreeNode{Val: 1, Right: node2}
	return root
}
func getTestData2() *TreeNode {
	root := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	node7 := &TreeNode{Val: 7}
	node8 := &TreeNode{Val: 8}
	node9 := &TreeNode{Val: 9}

	root.Left = node2
	root.Right = node3
	node2.Left = node4
	node2.Right = node5
	node5.Left = node6
	node3.Right = node8
	node8.Left = node7
	node8.Right = node9

	return root
}

/*
       1
      / \
     2   3
    / \   \
   4   5   8
      /   / \
     6   7   9
*/
