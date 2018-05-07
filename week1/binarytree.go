package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left

		invertTree(root.Left)
		invertTree(root.Right)
	}

	return root
}

func main() {

	root := TreeNode{0, new(TreeNode), new(TreeNode)}
	fmt.Println(root)
	invertTree(&root)

}
