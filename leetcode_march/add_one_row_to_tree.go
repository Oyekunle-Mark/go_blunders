package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	// when there is no depth d - 1
	if d == 1 {
		return &TreeNode{
			Val:   v,
			Left:  root,
			Right: nil,
		}
	}

	return root
}

func addNodeAtLevel(root *TreeNode, v, d, currentLevel int) {
	if currentLevel == d-1 {
		if root.Left != nil {
			root.Left = &TreeNode{
				Val:   v,
				Left:  root.Left,
				Right: nil,
			}
		}

		if root.Right != nil {
			root.Right = &TreeNode{
				Val:   v,
				Left:  nil,
				Right: root.Right,
			}
		}
	}

	addNodeAtLevel(root.Left, v, d, currentLevel+1)
	addNodeAtLevel(root.Right, v, d, currentLevel+1)
}
