package main

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	var dfs func(i, j int, grid [][]byte)

	dfs = func(i, j int, grid [][]byte) {
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}

		if grid[i][j] == '0' {
			return
		}

		grid[i][j] = '0'
		dfs(i-1, j, grid)
		dfs(i, j-1, grid)
		dfs(i+1, j, grid)
		dfs(i, j+1, grid)
	}

	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j, grid)
			}
		}
	}
	return res
}
