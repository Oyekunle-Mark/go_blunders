package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {

}

func inOrderNodes(node *TreeNode, current int, target int, nodeVals []int) {
	if node == nil {
		return
	}

	if current == target {
		nodeVals = append(nodeVals, node.Val)
		return
	}

	inOrderNodes(node.Left, current + 1, target, nodeVals)
	inOrderNodes(node.Right, current + 1, target, nodeVals)
}
