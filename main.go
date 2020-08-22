package main

import (
	"./solution"
	"./tree"
	"fmt"
)

func main() {
	quickTest([]int{7, 8, 0, 0, 12, 0, 0})

	quickTest([]int{9, 3, 0, 0, 4, 0, 0})

	quickTest([]int{3, 7, 8, 0, 0, 12, 0, 0, 9, 3, 0, 0, 4, 0, 0})

	fmt.Println("COMPLETE")
}

func quickTest(preorder []int) {
	root := tree.PreorderCreate(preorder)
	root.PreorderPrint()
	cost, depth := solution.MinCostToMaxDepthLeaf(root)
	fmt.Printf("cost: %d, depth: %d\n\n", cost, depth)
}
