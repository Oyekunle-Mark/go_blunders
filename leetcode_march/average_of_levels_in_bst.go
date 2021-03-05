package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	treeHeight := findTreeHeight(root, 0) // find tree height
	var averageNodePerLevel []float64     // to hold the averages

	// loop for every level int the tree
	for i := 1; i <= treeHeight; i++ {
		var nodesAtCurrentLevel []int                                                       // will hold current node val level i
		levelOrderNodes(root, 1, i, &nodesAtCurrentLevel)                                   // write node val to nodesAtCurrentLevel
		averageNodePerLevel = append(averageNodePerLevel, findAverage(nodesAtCurrentLevel)) // append the average to averageNodePerLevel
	}

	return averageNodePerLevel // return result
}

// levelOrderNodes appends the Val of all node at level target to nodeVals as side effect
func levelOrderNodes(node *TreeNode, current int, target int, nodeVals *[]int) {
	if node == nil {
		return
	}

	if current == target {
		*nodeVals = append(*nodeVals, node.Val)
		return
	}

	levelOrderNodes(node.Left, current+1, target, nodeVals)
	levelOrderNodes(node.Right, current+1, target, nodeVals)
}

// findAverage finds the average on the input slice of int and returns the result as a float64
func findAverage(nodeVals []int) float64 {
	var sum float64 = 0

	for _, val := range nodeVals {
		sum += float64(val)
	}

	return sum / float64(len(nodeVals))
}

// findTreeHeight finds the height of a bst.
func findTreeHeight(root *TreeNode, level int) int {
	if root == nil {
		return level
	}

	leftSubTreeHeight := findTreeHeight(root.Left, level+1)
	rightSubTreeHeight := findTreeHeight(root.Right, level+1)

	if leftSubTreeHeight >= rightSubTreeHeight {
		return leftSubTreeHeight
	} else {
		return rightSubTreeHeight
	}
}
