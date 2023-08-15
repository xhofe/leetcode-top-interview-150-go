package main

import "sort"

func findWords(board [][]byte, words []string) []string {
	tree := Constructor()
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for _, v := range words {
		tree.Insert(v)
	}
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	ans := make([]string, 0)
	var dfs func(i, j int, pre *TrieNode)
	dfs = func(i, j int, pre *TrieNode) {
		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] {
			return
		}
		b := board[i][j]
		idx := int(b - 'a')
		cur := pre.Children[idx]
		if cur == nil {
			return
		}
		if cur.IsEnd {
			ans = append(ans, cur.Word)
			cur.IsEnd = false
		}
		visited[i][j] = true
		for _, dir := range dirs {
			dfs(i+dir[0], j+dir[1], cur)
		}
		visited[i][j] = false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, tree.root)
		}
	}
	sort.Strings(ans)
	return ans
}
