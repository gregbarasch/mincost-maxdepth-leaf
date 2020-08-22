package tree

import "fmt"

type Node struct {
	Cost  int
	Left  *Node
	Right *Node
}

/**
Top level print function wrapper
*/
func (n *Node) PreorderPrint() {
	n.preorderPrint()
	fmt.Println()
}

/**
Recursive helper print function
*/
func (n *Node) preorderPrint() {
	fmt.Printf("%d ", n.Cost)

	if n.Left != nil {
		n.Left.preorderPrint()
	}

	if n.Right != nil {
		n.Right.preorderPrint()
	}
}

/**
@return a pointer to the root Node of the created tree

This function makes assumptions about elements of val 0
We use this to represent a termination condition within the tree. if a node has value 0, we assume this node nil
We do this to ensure that unique trees are create from a single preorder traversal
*/
func PreorderCreate(costs []int) *Node {
	root, _ := preorderCreate(costs)
	return root
}

/**
@return a pointer to the root Node, the remaining costs to be evaluated
Helper function that keeps track of "costs remaining" as an int slice
This function is recursively called
*/
func preorderCreate(costs []int) (*Node, []int) {

	// null/empty/0 costs
	if costs == nil || len(costs) == 0 {
		return nil, nil
	} else if costs[0] == 0 {
		return nil, costs[1:]
	}

	// Create current node
	self := &Node{
		Cost: costs[0],
	}

	// Left children recursively created using passed-in costs slice
	node, remaining := preorderCreate(costs[1:])
	self.Left = node

	// Right children recursively created using remaining elements of slice
	node, remaining = preorderCreate(remaining)
	self.Right = node

	// return self + remaining elements to be evaluated
	return self, remaining
}
