package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {

}

func levelOrderNodes(node *TreeNode, current int, target int, nodeVals []int) {
	if node == nil {
		return
	}

	if current == target {
		nodeVals = append(nodeVals, node.Val)
		return
	}

	levelOrderNodes(node.Left, current + 1, target, nodeVals)
	levelOrderNodes(node.Right, current + 1, target, nodeVals)
}
