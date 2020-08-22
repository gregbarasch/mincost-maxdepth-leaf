package solution

import (
	"../tree"
)

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}

/**
@return cost, depth

Uses a modified depth first search
*/
func MinCostToMaxDepthLeaf(root *tree.Node) (int, int) {

	// Depth of -1 for nil node.. Can't be zero, because will conflict with single node trees
	if root == nil {
		return 0, -1
	}

	// Leaf node
	if root.Left == nil && root.Right == nil {
		return root.Cost, 0
	}

	lcost, ldepth := MinCostToMaxDepthLeaf(root.Left)
	rcost, rdepth := MinCostToMaxDepthLeaf(root.Right)

	if ldepth > rdepth {
		return root.Cost + lcost, 1 + ldepth
	} else if ldepth == rdepth {
		return root.Cost + min(lcost, rcost), 1 + ldepth
	} else {
		return root.Cost + rcost, 1 + rdepth
	}
}
