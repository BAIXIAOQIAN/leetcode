package main

type TrieNode struct {
	Word   string
	Chirld [26]*TrieNode
}

func findWords(board [][]byte, words []string) []string {
	root := &TrieNode{}

	//初始化字典树
	for _, v := range words {
		node := root
		for _, w := range v {
			if node.Chirld[w-'a'] == nil {
				node.Chirld[w-'a'] = &TrieNode{}
			}
			node = node.Chirld[w-'a']
		}
		node.Word = v
	}

	res := make([]string, 0)
	m, n := len(board), len(board[0])

	var dfs func(i, j int, node *TrieNode)
	dfs = func(i, j int, node *TrieNode) {
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}

		c := board[i][j]
		if c == '#' || node.Chirld[c-'a'] == nil {
			return
		}
		node = node.Chirld[c-'a']

		if node.Word != "" {
			res = append(res, node.Word)
			node.Word = ""
		}

		board[i][j] = '#'
		dfs(i-1, j, node)
		dfs(i+1, j, node)
		dfs(i, j-1, node)
		dfs(i, j+1, node)
		board[i][j] = c
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, root)
		}
	}
	return res
}
