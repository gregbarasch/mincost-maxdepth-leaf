package solution

import (
	"../tree"
	"testing"
)

func TestMinCostToMaxDepthLeaf(t *testing.T) {

	tests := []struct {
		name      string
		givenRoot *tree.Node
		expCost   int
		expDepth  int
	}{
		{
			name:      "Left Sided Linked List",
			givenRoot: tree.PreorderCreate([]int{3, 4, 5, 6, 10, 0, 0, 0, 0}),
			expCost:   28,
			expDepth:  4,
		},
		{
			name:      "Right Sided Linked List",
			givenRoot: tree.PreorderCreate([]int{8, 0, 7, 0, 9, 0, 3, 0, 12, 0, 10, 0, 0}),
			expCost:   49,
			expDepth:  5,
		},
		{
			name:      "Perfectly Balanced Tree",
			givenRoot: tree.PreorderCreate([]int{3, 7, 8, 0, 0, 12, 0, 0, 9, 3, 0, 0, 4, 0, 0}),
			expCost:   15,
			expDepth:  2,
		},
		{
			name:      "Unbalanced Tree",
			givenRoot: tree.PreorderCreate([]int{17, 3, 0, 0, 5, 7, 0, 0, 9, 0, 0}),
			expCost:   29,
			expDepth:  2,
		},
		{
			name:      "Single Node",
			givenRoot: tree.PreorderCreate([]int{25}),
			expCost:   25,
			expDepth:  0,
		},
		{
			name:      "Nil",
			givenRoot: nil,
			expCost:   0,
			expDepth:  -1,
		},
		{
			name:      "Negative Numbers Included",
			givenRoot: tree.PreorderCreate([]int{-10, -5, 0, 0, -22, -51, 0, 5, 0, 0, -100, 5, 0, 0, 10, 0, 0}),
			expCost:   -127,
			expDepth:  3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			cost, depth := MinCostToMaxDepthLeaf(test.givenRoot)

			if test.expCost != cost {
				t.Errorf("Cost %d does not match expected cost %d", cost, test.expCost)
			}

			if test.expDepth != depth {
				t.Errorf("Depth %d does not match expected depth %d", depth, test.expDepth)
			}
		})
	}
}
