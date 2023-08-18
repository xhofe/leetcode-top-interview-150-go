package main

func isValidBST(root *TreeNode) bool {
	prev := 0
	first := true
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !dfs(node.Left) {
			return false
		}
		if first {
			prev = node.Val
			first = false
		} else {
			if node.Val <= prev {
				return false
			}
			prev = node.Val
		}
		return dfs(node.Right)
	}
	return dfs(root)
}
