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

	levelOrderNodes(node.Left, current+1, target, nodeVals)
	levelOrderNodes(node.Right, current+1, target, nodeVals)
}

func findAverage(nodeVals []int) float64 {
	var sum float64 = 0

	for _, val := range nodeVals {
		sum += float64(val)
	}

	return sum / float64(len(nodeVals))
}

func findTreeHeight(root *TreeNode, level int) int {
	if root == nil {
		return level
	}

	leftSubTree := findTreeHeight(root.Left, level + 1)
	rightSubTree := findTreeHeight(root.Right, level + 1)

	if leftSubTree >= rightSubTree {
		return leftSubTree
	} else {
		return rightSubTree
	}
}
