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

	addNodeAtLevel(root, v, d, 1)

	return root
}

// addNodeAtLevel adds a node at the appropriate level.
// For each NOT null tree nodes N in depth d-1,
// create two tree nodes with value v as N's left subtree root and right subtree root.
// And N's original left subtree should be the left subtree of the new left subtree root,
// its original right subtree should be the right subtree of the new right subtree root.
// For each null tree nodes N in depth d - 1, creat a tree node with value v.
func addNodeAtLevel(root *TreeNode, v, d, currentLevel int) {
	// return when root is nil
	if root == nil {
		return
	}

	// at level d - 1, i.e one above desired level
	if currentLevel == d-1 {
		root.Left = &TreeNode{
			Val:   v,
			Left:  root.Left,
			Right: nil,
		}

		root.Right = &TreeNode{
			Val:   v,
			Left:  nil,
			Right: root.Right,
		}

		return
	}

	addNodeAtLevel(root.Left, v, d, currentLevel+1)  // recursively process the left sub tree
	addNodeAtLevel(root.Right, v, d, currentLevel+1) // recursively process the left sub tree
}
