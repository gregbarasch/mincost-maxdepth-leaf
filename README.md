Given a binary tree, find the minimum cost required to reach a leaf node at the deepest depth of said tree.

The following is an example definition of a node:
```
struct node {
    int cost;
    struct node *left;
    struct node *right;
};
```
<br>

The final cost will be a sum of the costs of all nodes along the path to the leaf, including the costs of both the root and leaf node.<br>

The function might look something like this:
```
int calcMinCostToMaxDepthLeaf(node *root) {
    return totalCost;
}
```