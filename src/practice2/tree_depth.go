package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	maxValue := 0
	for _, child := range root.Children {
		if childDepth := maxDepth(child); childDepth > maxValue {
			maxValue = childDepth
		}
	}
	return maxValue + 1
}
